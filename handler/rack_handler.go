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
}

func NewRackHandler(s service.Service) RackHandler {
	return RackHandler{Service: s}
}

func (h *RackHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	racks, totalRecords, totalPages, err := h.Service.RackService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "Failed to get racks", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of racks", http.StatusOK, racks, &utils.Pagination{
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
		utils.WriteError(w, "Rack not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Rack found", http.StatusOK, Rack, nil)
}

func (h *RackHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Service.RackService.Create(req)
	if err != nil {
		utils.WriteError(w, "Failed to create Rack", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Rack created", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *RackHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateRackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.RackService.Update(id, req); err != nil {
		utils.WriteError(w, "Failed to update Rack", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Rack updated", http.StatusOK, nil, nil)
}

func (h *RackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.RackService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete Rack", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Rack deleted", http.StatusOK, nil, nil)
}
