package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type AddressRepository struct {
	pool *pgxpool.Pool
}

func NewAddressRepository(pool *pgxpool.Pool) *AddressRepository {
	return &AddressRepository{pool: pool}
}

func (r *AddressRepository) ListByUserID(ctx context.Context, userID string) ([]domain.Address, error) {
	const query = `
		SELECT id, user_id, recipient_name, phone_number, province, city,
		       district, postal_code, full_address, COALESCE(label, ''),
		       COALESCE(coordinates, ''), is_primary, 
		       COALESCE(created_at, NOW()), COALESCE(updated_at, NOW())
		FROM addresses
		WHERE user_id = $1
		ORDER BY is_primary DESC, created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("address: list by user id: %w", err)
	}
	defer rows.Close()

	var addresses []domain.Address
	for rows.Next() {
		var addr domain.Address
		if err := rows.Scan(
			&addr.ID, &addr.UserID, &addr.RecipientName, &addr.PhoneNumber,
			&addr.Province, &addr.City, &addr.District, &addr.PostalCode,
			&addr.FullAddress, &addr.Label, &addr.Coordinates, &addr.IsPrimary,
			&addr.CreatedAt, &addr.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("address: scan row: %w", err)
		}
		addresses = append(addresses, addr)
	}

	return addresses, nil
}

func (r *AddressRepository) Create(ctx context.Context, address domain.Address) (domain.Address, error) {
	const query = `
		INSERT INTO addresses (user_id, recipient_name, phone_number, province, city, district, postal_code, full_address, label, is_primary)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NULLIF($9, ''), $10)
		RETURNING id, user_id, recipient_name, phone_number, province, city,
		          district, postal_code, full_address, COALESCE(label, ''),
		          COALESCE(coordinates, ''), is_primary, 
		          COALESCE(created_at, NOW()), COALESCE(updated_at, NOW())
	`

	var addr domain.Address
	err := r.pool.QueryRow(ctx, query,
		address.UserID, address.RecipientName, address.PhoneNumber,
		address.Province, address.City, address.District, address.PostalCode,
		address.FullAddress, address.Label, address.IsPrimary,
	).Scan(
		&addr.ID, &addr.UserID, &addr.RecipientName, &addr.PhoneNumber,
		&addr.Province, &addr.City, &addr.District, &addr.PostalCode,
		&addr.FullAddress, &addr.Label, &addr.Coordinates, &addr.IsPrimary,
		&addr.CreatedAt, &addr.UpdatedAt,
	)
	if err != nil {
		return domain.Address{}, fmt.Errorf("address: create: %w", err)
	}

	return addr, nil
}

func (r *AddressRepository) Delete(ctx context.Context, id string, userID string) error {
	const query = `DELETE FROM addresses WHERE id = $1 AND user_id = $2`

	result, err := r.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("address: delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("address not found")
	}

	return nil
}

func (r *AddressRepository) SetPrimary(ctx context.Context, id string, userID string) error {
	// Unique partial index di DB memastikan hanya satu is_primary = true per user.
	// Kita update dalam satu transaksi: reset semua, set yang dipilih.
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("address: set primary: begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`UPDATE addresses SET is_primary = false WHERE user_id = $1`, userID,
	); err != nil {
		return fmt.Errorf("address: set primary: reset: %w", err)
	}

	result, err := tx.Exec(ctx,
		`UPDATE addresses SET is_primary = true WHERE id = $1 AND user_id = $2`, id, userID,
	)
	if err != nil {
		return fmt.Errorf("address: set primary: update: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("address not found")
	}

	return tx.Commit(ctx)
}
