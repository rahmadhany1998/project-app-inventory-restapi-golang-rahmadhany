package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/service"
	"project-app-inventory-restapi-golang-rahmadhany/utils"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	Service service.Service
	Config  utils.Configuration
}

func NewProductHandler(s service.Service, cfg utils.Configuration) ProductHandler {
	return ProductHandler{Service: s, Config: cfg}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit := h.Config.Limit
	products, totalRecords, totalPages, err := h.Service.ProductService.GetAll(page, limit)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, products, &utils.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	})
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	product, err := h.Service.ProductService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}
	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, product, nil)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	// batas maksimal file upload (10 MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteError(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Ambil field dari form-data
	name := r.FormValue("name")
	categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
	rackID, _ := strconv.Atoi(r.FormValue("rack_id"))
	warehouseID, _ := strconv.Atoi(r.FormValue("warehouse_id"))
	inventoryCount, _ := strconv.Atoi(r.FormValue("inventory_count"))
	retailPrice, _ := strconv.Atoi(r.FormValue("retail_price"))
	sellingPrice, _ := strconv.Atoi(r.FormValue("selling_price"))

	// Ambil file
	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, "Image is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Simpan file ke folder lokal
	filename := fmt.Sprintf("static/uploads/%d_%s", time.Now().Unix(), handler.Filename)
	savepath := fmt.Sprintf("/static/uploads/%d_%s", time.Now().Unix(), handler.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}

	// Simpan data produk ke DB
	product := dto.CreateProductRequest{
		Name:           name,
		CategoryID:     categoryID,
		RackID:         rackID,
		WarehouseID:    warehouseID,
		InventoryCount: inventoryCount,
		RetailPrice:    retailPrice,
		SellingPrice:   sellingPrice,
		Image:          savepath, // path ke file yang disimpan
	}

	validation, err := utils.ValidateData(product)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	id, err := h.Service.ProductService.Create(product)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccess(w, "Data created successfully", http.StatusCreated, map[string]int{"id": id}, nil)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Ambil ID produk dari URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // batas 10MB
	if err != nil {
		utils.WriteError(w, "Failed to parse form-data", http.StatusBadRequest)
		return
	}

	// Ambil data lama dari DB
	existing, err := h.Service.ProductService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}

	// Ambil field dari form-data
	name := r.FormValue("name")
	categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
	rackID, _ := strconv.Atoi(r.FormValue("rack_id"))
	warehouseID, _ := strconv.Atoi(r.FormValue("warehouse_id"))
	inventoryCount, _ := strconv.Atoi(r.FormValue("inventory_count"))
	retailPrice, _ := strconv.Atoi(r.FormValue("retail_price"))
	sellingPrice, _ := strconv.Atoi(r.FormValue("selling_price"))

	// Cek apakah user upload file baru
	var imagePath, savepath string
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		imagePath = fmt.Sprintf("static/uploads/%d_%s", time.Now().Unix(), handler.Filename)
		savepath = fmt.Sprintf("/static/uploads/%d_%s", time.Now().Unix(), handler.Filename)

		dst, err := os.Create(imagePath)
		if err != nil {
			utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
			return
		}
		if existing.Image != "" {
			path := strings.TrimPrefix(existing.Image, "/")
			_ = os.Remove(path)
		}

	} else {
		// Tidak upload file → pakai image sebelumnya
		savepath = existing.Image
	}

	// Update object
	product := dto.UpdateProductRequest{
		Name:           name,
		CategoryID:     categoryID,
		RackID:         rackID,
		WarehouseID:    warehouseID,
		InventoryCount: inventoryCount,
		RetailPrice:    retailPrice,
		SellingPrice:   sellingPrice,
		Image:          savepath,
	}

	validation, err := utils.ValidateData(product)
	if err != nil {
		utils.ResponseErrorValidation(w, http.StatusBadRequest, "Validation error", validation)
		return
	}

	err = h.Service.ProductService.Update(id, product)
	if err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccess(w, "Data processed successfully.", http.StatusOK, nil, nil)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	existing, err := h.Service.ProductService.GetByID(id)
	if err != nil {
		utils.WriteError(w, "Data not found", http.StatusNotFound)
		return
	}

	path := strings.TrimPrefix(existing.Image, "/")
	_ = os.Remove(path)

	if err := h.Service.ProductService.Delete(id); err != nil {
		utils.WriteError(w, "An internal server error occurred.", http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w, "Data deleted", http.StatusOK, nil, nil)
}
