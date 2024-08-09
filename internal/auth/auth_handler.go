package auth

import (
	"net/http"

	"github.com/aslam-ep/go-e-commerce/utils"
)

// Handler handles HTTP requests related to authentication.
type Handler struct {
	service Service
}

// NewHandler creates a new instance of the Handler with the provided authentication service.
func NewHandler(as Service) *Handler {
	return &Handler{
		service: as,
	}
}

// Register      godoc
// @Summary      Register a new user
// @Description  Register a new user with the provided details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  RegisterUserReq  true  "User registration request"
// @Success      200  {object}  user.User
// @Failure      400  {object}  utils.MessageRes
// @Router       /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var registerUserReq RegisterUserReq
	if err := utils.ReadFromRequest(r, &registerUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(registerUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.RegisterUser(r.Context(), &registerUserReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// Login         godoc
// @Summary      Login user
// @Description  Login a user, on success get the refreshToken and accessToken
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  LoginReq  true  "Login request"
// @Success      200  {object}  LoginRes "Login response"
// @Failure      400  {object}  utils.MessageRes "Default response"
// @Failure      401  {object}  utils.MessageRes "Default response"
// @Router       /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	if err := utils.ReadFromRequest(r, &req); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Authenticate(r.Context(), &req)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusAccepted, res)
}

// RefreshToken  godoc
// @Summary      Refresh token
// @Description  Refresh token, send the new access token based on refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  RefreshTokenReq  true  "Refresh token request"
// @Success      200  {object}  RefreshTokenRes "Refresh token response"
// @Failure      400  {object}  utils.MessageRes "Default response"
// @Failure      401  {object}  utils.MessageRes "Default response"
// @Router       /auth/refresh-token [post]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenReq
	if err := utils.ReadFromRequest(r, &req); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.RefreshToken(r.Context(), &req)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusAccepted, res)
}
