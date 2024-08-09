package user

import (
	"net/http"
	"strconv"

	"github.com/aslam-ep/go-e-commerce/utils"
	"github.com/go-chi/chi/v5"
)

// Handler struct to hold the user service and provide handler functions
type Handler struct {
	service Service
}

// NewHandler initialize and return the user Handler
func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

// GetUser       godoc
// @Summary      Get User Details
// @Description  Get User Details by provided ID in url
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{user_id} [post]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// UpdateUser    godoc
// @Summary      Update User Details
// @Description  Update User Details by provided ID in url and details in body
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Param        body  body  UpdateUserReq  true  "User Update request"
// @Success      200  {object}  User
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{user_id}/update [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var updateUserReq UpdateUserReq
	if err := utils.ReadFromRequest(r, &updateUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	updateUserReq.ID = int64(userID)

	if err := utils.Validate.Struct(updateUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.UpdateUser(r.Context(), &updateUserReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// ChangePassword godoc
// @Summary      Reset User Password
// @Description  Reset User Password by provided ID in url and password in body
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Param        body  body  ResetPasswordReq  true  "Password change request"
// @Success      200  {object}  utils.MessageRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{user_id}/password-reset [put]
func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var resetPasswordReq ResetPasswordReq
	if err := utils.ReadFromRequest(r, &resetPasswordReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	resetPasswordReq.ID = int64(userID)

	res, err := h.service.ChangeUserPassword(r.Context(), &resetPasswordReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// DeleteUser    godoc
// @Summary      Delete User Details
// @Description  Delete User Details by provided ID in url
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  utils.MessageRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{user_id}/delete [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.DeleteUser(r.Context(), userID)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}
