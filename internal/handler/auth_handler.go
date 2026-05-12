package handler

import (
	"net/http"
	"time"

	"github.com/thisgleammm/mantis-backend/internal/domain"
	"github.com/thisgleammm/mantis-backend/internal/json"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login handles user authentication.
// @Summary Login
// @Description Authenticate user and return access/refresh tokens in cookies.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body loginRequest true "Login request"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {string} string "unauthorized"
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, accessToken, refreshToken, err := h.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/api/v1/auth/refresh",
		SameSite: http.SameSiteNoneMode,
	})

	json.Write(w, http.StatusOK, map[string]any{
		"user":  user,
		"token": accessToken,
	})
}

type registerRequest struct {
	Username    string `json:"username" validate:"required,min=3,max=30"`
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

// Register handles user registration.
// @Summary Register
// @Description Create a new user account.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body registerRequest true "Register request"
// @Success 201 {object} domain.User
// @Failure 400 {string} string "invalid request"
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user := domain.User{
		Username:    req.Username,
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}

	created, err := h.svc.Register(r.Context(), user)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusCreated, created)
}

// Refresh handles token refreshment.
// @Summary Refresh Token
// @Description Refresh access token using refresh token cookie.
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "unauthorized"
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, "missing refresh token")
		return
	}

	accessToken, refreshToken, err := h.svc.RefreshToken(r.Context(), cookie.Value)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/api/v1/auth/refresh",
		SameSite: http.SameSiteNoneMode,
	})

	json.Write(w, http.StatusOK, map[string]any{
		"token": accessToken,
	})
}

// Logout handles user logout.
// @Summary Logout
// @Description Clear authentication cookies.
// @Tags Auth
// @Success 204
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		Path:     "/api/v1/auth/refresh",
		SameSite: http.SameSiteNoneMode,
	})
	w.WriteHeader(http.StatusNoContent)
}

type forgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ForgotPassword handles password reset request.
// @Summary Forgot Password
// @Description Request a password reset email.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body forgotPasswordRequest true "Forgot password request"
// @Success 200 {object} map[string]string
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req forgotPasswordRequest
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := h.svc.ForgotPassword(r.Context(), req.Email); err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "password reset link sent"})
}

type resetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// ResetPassword handles password reset.
// @Summary Reset Password
// @Description Reset password using a valid token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body resetPasswordRequest true "Reset password request"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "invalid or expired token"
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req resetPasswordRequest
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := h.svc.ResetPassword(r.Context(), req.Token, req.Password); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "password reset successful"})
}

type confirmPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

// ConfirmPassword handles password confirmation for sensitive actions.
// @Summary Confirm Password
// @Description Verify current user password.
// @Tags Auth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body confirmPasswordRequest true "Confirm password request"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "unauthorized"
// @Router /auth/confirm-password [post]
func (h *AuthHandler) ConfirmPassword(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req confirmPasswordRequest
	if err := json.Read(w, r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := h.svc.ConfirmPassword(r.Context(), userID, req.Password); err != nil {
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "password confirmed"})
}
