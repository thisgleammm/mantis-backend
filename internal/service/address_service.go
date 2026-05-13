package service

import (
	"context"

	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type AddressService struct {
	repo domain.AddressRepository
}

func NewAddressService(repo domain.AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) ListAddresses(ctx context.Context, userID string) ([]domain.Address, error) {
	return s.repo.ListByUserID(ctx, userID)
}

func (s *AddressService) CreateAddress(ctx context.Context, address domain.Address) (domain.Address, error) {
	return s.repo.Create(ctx, address)
}

func (s *AddressService) DeleteAddress(ctx context.Context, id string, userID string) error {
	return s.repo.Delete(ctx, id, userID)
}

func (s *AddressService) SetPrimaryAddress(ctx context.Context, id string, userID string) error {
	return s.repo.SetPrimary(ctx, id, userID)
}
