package product

import (
	"app/internal/domain/product"
	productService "app/internal/service/product"
	"app/internal/utils/validation"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *productService.Service
}

func NewHandler(service *productService.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input product.CreateProductRequest

	validationErrors := validation.StructValidator(&input)

	if len(validationErrors) != 0 {
		w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(product)
}

// func (h *ProductHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
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

// func (h *ProductHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
// 	page := 1
// 	if v, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && v > 0 {
// 		page = v
// 	}

// 	categories, p, err := h.service.FindAll(page)

// 	if err != nil {
// 		http.Error(w, "category not found", http.StatusNotFound)
// 		return
// 	}

// 	response := map[string]interface{}{
// 		"data": categories,
// 		"meta": p,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	_ = json.NewEncoder(w).Encode(response)
// }
