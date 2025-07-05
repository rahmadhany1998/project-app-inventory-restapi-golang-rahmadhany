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

type SaleHandler struct {
	Service service.Service
	Config  utils.Configuration
}

func NewSaleHandler(s service.Service, cfg utils.Configuration) SaleHandler {
	return SaleHandler{Service: s, Config: cfg}
}

func (h *SaleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := h.Config.Limit
	sales, totalRecords, totalPages, err := h.Service.SaleService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of sales", http.StatusOK, sales, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *SaleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	sale, err := h.Service.SaleService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, sale, nil)
}

func (h *SaleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validation, err := utils.ValidateData(req)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	id, err := h.Service.SaleService.Create(req)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data created successfully", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *SaleHandler) GetReportSummaryByDate(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start_date")
	end := r.URL.Query().Get("end_date")

	if start == "" || end == "" {
		utils.WriteError(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	report, err := h.Service.SaleService.GetReportSummaryByDate(start, end)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, report, nil)
}
