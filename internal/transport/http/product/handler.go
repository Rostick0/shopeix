package product

import (
	"app/internal/domain/product"
	productService "app/internal/service/product"
	"app/internal/utils/validation"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *productService.Service
}

func NewHandler(service *productService.Service) *Handler {
	return &Handler{service: service}
}

// func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
// 	idParam := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idParam, 10, 64)
// 	if err != nil {
// 		http.Error(w, "invalid category id", http.StatusBadRequest)
// 		return
// 	}

// 	u, err := h.service.FindByID(id)
// 	if err != nil {
// 		http.Error(w, "category not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	_ = json.NewEncoder(w).Encode(u)
// }

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	// var input product.GetListProductRequest

	// r.URL.Query().
	// r.Context().Value(r.URL.Query())(&input)
	// r.URL.Parse()
	// r.URL.Query().Encode()
	// input:=r.URL.Query()
	// json.NewDecoder(r.URL.Query()).Decode(&input)

	page := 1
	if v, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && v > 0 {
		page = v
	}

	products, p, err := h.service.FindAll(page)

	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"data": products,
		"meta": p,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input product.CreateProductRequest

	json.NewDecoder(r.Body).Decode(&input)
	validationErrors := validation.StructValidator(&input)
	w.Header().Set("Content-Type", "application/json")

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	product, err := h.service.Create(&input)

	if err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(product)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	productFinded, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	var input product.UpdateProductRequest

	json.NewDecoder(r.Body).Decode(&input)
	validationErrors := validation.StructValidator(&input)
	w.Header().Set("Content-Type", "application/json")

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	product, err := h.service.Update(&input, productFinded)

	if err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(product)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	productFinded, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	product, err := h.service.Delete(productFinded)

	if err != nil {
		http.Error(w, "failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(product)
}
