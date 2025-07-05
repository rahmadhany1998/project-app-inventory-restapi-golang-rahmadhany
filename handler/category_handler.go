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
	Config  utils.Configuration
}

func NewCategoryHandler(s service.Service, cfg utils.Configuration) CategoryHandler {
	return CategoryHandler{
		Service: s,
		Config:  cfg,
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := h.Config.Limit

	categories, totalRecords, totalPages, err := h.Service.CategoryService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data Processed Successfully", http.StatusOK, categories, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category, err := h.Service.CategoryService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Data Processed Successfully", http.StatusOK, category, nil)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	id, err := h.Service.CategoryService.Create(req)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data Created Successfully", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	if err := h.Service.CategoryService.Update(id, req); err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data Processed Successfully", http.StatusOK, nil, nil)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.CategoryService.Delete(id); err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data deleted", http.StatusNoContent, nil, nil)
}
