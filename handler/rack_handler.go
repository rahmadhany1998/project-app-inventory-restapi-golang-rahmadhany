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

type RackHandler struct {
	Service service.Service
	Config  utils.Configuration
}

func NewRackHandler(s service.Service, cfg utils.Configuration) RackHandler {
	return RackHandler{Service: s, Config: cfg}
}

func (h *RackHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := h.Config.Limit
	racks, totalRecords, totalPages, err := h.Service.RackService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, racks, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *RackHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	Rack, err := h.Service.RackService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, Rack, nil)
}

func (h *RackHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	id, err := h.Service.RackService.Create(req)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data Created successfully.", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *RackHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateRackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	if err := h.Service.RackService.Update(id, req); err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, nil, nil)
}

func (h *RackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.RackService.Delete(id); err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Rack deleted", http.StatusNoContent, nil, nil)
}
