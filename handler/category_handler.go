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

type CategoryHandler struct {
	Service service.Service
}

func NewCategoryHandler(s service.Service) CategoryHandler {
	return CategoryHandler{Service: s}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.CategoryService.GetAll()
	if err != nil {
		utils.WriteError(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of categories", categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category, err := h.Service.CategoryService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Category not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Category found", category)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Service.CategoryService.Create(req)
	if err != nil {
		utils.WriteError(w, "Failed to create category", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Category created", map[string]int{"id": id})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.CategoryService.Update(id, req); err != nil {
		utils.WriteError(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Category updated", nil)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.CategoryService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Category deleted", nil)
}
