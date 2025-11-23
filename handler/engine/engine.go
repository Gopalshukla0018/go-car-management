package engine

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/Gopalshukla0018/go-car-management/service/engine"
)

type Handler struct {
	service *engine.Service
}

func New(service *engine.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.EngineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	if err := models.ValidateEngineRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateEngine(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 { 
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return 
	}
	id := parts[len(parts)-1]

	resp, err := h.service.GetEngine(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}