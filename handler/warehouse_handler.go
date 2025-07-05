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

type WarehouseHandler struct {
	Service service.Service
}

func NewWarehouseHandler(s service.Service) WarehouseHandler {
	return WarehouseHandler{Service: s}
}

func (h *WarehouseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	warehouses, totalRecords, totalPages, err := h.Service.WarehouseService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "Failed to get warehouses", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of warehouses", http.StatusOK, warehouses, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *WarehouseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	Warehouse, err := h.Service.WarehouseService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Warehouse not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Warehouse found", http.StatusOK, Warehouse, nil)
}

func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	id, err := h.Service.WarehouseService.Create(req)
	if err != nil {
		utils.WriteError(w, "Failed to create Warehouse", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Warehouse created", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *WarehouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	if err := h.Service.WarehouseService.Update(id, req); err != nil {
		utils.WriteError(w, "Failed to update Warehouse", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Warehouse updated", http.StatusOK, nil, nil)
}

func (h *WarehouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.WarehouseService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete Warehouse", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Warehouse deleted", http.StatusOK, nil, nil)
}
