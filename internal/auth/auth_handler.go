package auth

import (
	"net/http"

	"github.com/aslam-ep/go-e-commerce/utils"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(as AuthService) *AuthHandler {
	return &AuthHandler{
		authService: as,
	}
}

// @Summary      Login user
// @Description  Login a user, on success get the refreshToken and accessToken
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  LoginReq  true  "Login request"
// @Success      200  {object}  LoginRes "Login response"
// @Failure      400  {object}  utils.MessageRes "Default response"
// @Failure      401  {object}  utils.MessageRes "Default response"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	if err := utils.ReadFromRequest(r, &req); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.authService.Authenticate(r.Context(), &req)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusAccepted, res)
}

// @Summary      Refresh token
// @Description  Refresh token, send the new access token based on refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  RefreshTokenReq  true  "Refresh token request"
// @Success      200  {object}  RefreshTokenRes "Refresh token response"
// @Failure      400  {object}  utils.MessageRes "Default response"
// @Failure      401  {object}  utils.MessageRes "Default response"
// @Router       /auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenReq
	if err := utils.ReadFromRequest(r, &req); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.authService.RefreshToken(r.Context(), &req)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusAccepted, res)
}
