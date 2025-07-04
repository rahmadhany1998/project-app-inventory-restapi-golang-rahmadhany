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

type ProductHandler struct {
	Service service.Service
}

func NewProductHandler(s service.Service) ProductHandler {
	return ProductHandler{Service: s}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.ProductService.GetAll()
	if err != nil {
		utils.WriteError(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "List of products", products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	product, err := h.Service.ProductService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Product not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Product found", product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Service.ProductService.Create(req)
	if err != nil {
		utils.WriteError(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Product created", map[string]int{"id": id})
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.ProductService.Update(id, req); err != nil {
		utils.WriteError(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Product updated", nil)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.ProductService.Delete(id); err != nil {
		utils.WriteError(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Product deleted", nil)
}
