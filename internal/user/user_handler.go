package user

import (
	"net/http"
	"strconv"

	"github.com/aslam-ep/go-e-commerce/utils"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

// @Summary      Register a new user
// @Description  Register a new user with the provided details
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body  CreateUserReq  true  "User registration request"
// @Success      200  {object}  UserRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/register [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserReq CreateUserReq
	if err := utils.ReadFromRequest(r, &createUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(createUserReq); err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.Service.CreateUser(r.Context(), &createUserReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// @Summary      Get User Details
// @Description  Get User Details by provided ID in url
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  UserRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{id} [post]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.Service.GetUserById(r.Context(), userID)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// @Summary      Update User Details
// @Description  Update User Details by provided ID in url and details in body
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Param        body  body  UpdateUserReq  true  "User Update request"
// @Success      200  {object}  UserRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{id}/update [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "id")
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

	res, err := h.Service.UpdateUser(r.Context(), &updateUserReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// @Summary      Reset User Password
// @Description  Reset User Password by provided ID in url and password in body
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Param        body  body  ResetPasswordReq  true  "Password change request"
// @Success      200  {object}  utils.MessageRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{id}/password-reset [put]
func (h *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "id")
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

	res, err := h.Service.ResetUserPassword(r.Context(), &resetPasswordReq)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

// @Summary      Delete User Details
// @Description  Delete User Details by provided ID in url
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  utils.MessageRes
// @Failure      400  {object}  utils.MessageRes
// @Router       /users/{id}/delete [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.Service.DeleteUser(r.Context(), userID)
	if err != nil {
		utils.WriterErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}
