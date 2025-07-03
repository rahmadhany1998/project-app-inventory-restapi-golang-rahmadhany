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
	racks, err := h.Service.RackService.GetAll()
	if err != nil {
		utils.WriteError(w, "Failed to get racks", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of racks", racks)
}

func (h *RackHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	Rack, err := h.Service.RackService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Rack not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Rack found", Rack)
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
	utils.WriteSuccess(w, "Rack created", map[string]int{"id": id})
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
	utils.WriteSuccess(w, "Rack updated", nil)
}

func (h *RackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.RackService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete Rack", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Rack deleted", nil)
}
