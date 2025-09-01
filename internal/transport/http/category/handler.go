package category

import (
	categoryService "app/internal/service/category"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *categoryService.Service
}

func NewHandler(service *categoryService.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid category id", http.StatusBadRequest)
		return
	}

	u, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, "category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(u)
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	page := 1
	if v, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && v > 0 {
		page = v
	}

	categories, p, err := h.service.FindAll(page)

	if err != nil {
		http.Error(w, "category not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"data": categories,
		"meta": p,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
