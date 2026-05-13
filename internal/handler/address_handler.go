package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thisgleammm/mantis-backend/internal/domain"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type AddressHandler struct {
	svc *service.AddressService
}

func NewAddressHandler(svc *service.AddressService) *AddressHandler {
	return &AddressHandler{svc: svc}
}

// ListAddresses returns all addresses for the authenticated user.
// @Summary List Addresses
// @Description Get all saved addresses for the current user.
// @Tags Addresses
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} domain.Address
// @Failure 401 {string} string "unauthorized"
// @Router /addresses [get]
func (h *AddressHandler) ListAddresses(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	addresses, err := h.svc.ListAddresses(r.Context(), userID)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, addresses)
}

// CreateAddress saves a new address for the authenticated user.
// @Summary Create Address
// @Description Save a new delivery address.
// @Tags Addresses
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 201 {object} domain.Address
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Router /addresses [post]
func (h *AddressHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req struct {
		RecipientName string `json:"recipient_name"`
		PhoneNumber   string `json:"phone_number"`
		Province      string `json:"province"`
		City          string `json:"city"`
		District      string `json:"district"`
		PostalCode    string `json:"postal_code"`
		FullAddress   string `json:"full_address"`
		Label         string `json:"label"`
		IsPrimary     bool   `json:"is_primary"`
	}

	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.RecipientName == "" || req.PhoneNumber == "" || req.FullAddress == "" {
		json.WriteError(w, http.StatusBadRequest, "recipient_name, phone_number, and full_address are required")
		return
	}

	address := domain.Address{
		UserID:        userID,
		RecipientName: req.RecipientName,
		PhoneNumber:   req.PhoneNumber,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		PostalCode:    req.PostalCode,
		FullAddress:   req.FullAddress,
		Label:         req.Label,
		IsPrimary:     req.IsPrimary,
	}

	created, err := h.svc.CreateAddress(r.Context(), address)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusCreated, created)
}

// DeleteAddress removes a saved address.
// @Summary Delete Address
// @Description Delete one of the current user's saved addresses.
// @Tags Addresses
// @Security ApiKeyAuth
// @Param id path string true "Address ID"
// @Success 204
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Router /addresses/{id} [delete]
func (h *AddressHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id := chi.URLParam(r, "id")

	if err := h.svc.DeleteAddress(r.Context(), id, userID); err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SetPrimaryAddress marks an address as the primary delivery address.
// @Summary Set Primary Address
// @Description Mark one of the current user's addresses as primary.
// @Tags Addresses
// @Security ApiKeyAuth
// @Param id path string true "Address ID"
// @Success 204
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Router /addresses/{id}/primary [patch]
func (h *AddressHandler) SetPrimaryAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id := chi.URLParam(r, "id")

	if err := h.svc.SetPrimaryAddress(r.Context(), id, userID); err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
