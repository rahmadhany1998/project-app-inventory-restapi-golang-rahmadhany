package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/service"
	"project-app-inventory-restapi-golang-rahmadhany/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service service.Service
}

func NewUserHandler(s service.Service) UserHandler {
	return UserHandler{Service: s}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.UserService.GetAll()
	if err != nil {
		utils.WriteError(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of users", users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.Service.UserService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "User not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "User found", user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Service.UserService.Create(req)
	if err != nil {
		utils.WriteError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "User created", map[string]int{"id": id})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.UserService.Update(id, req); err != nil {
		utils.WriteError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "User updated", nil)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.UserService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "User deleted", nil)
}
