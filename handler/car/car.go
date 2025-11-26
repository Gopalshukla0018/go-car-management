package car

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/Gopalshukla0018/go-car-management/service/car"
)

type Handler struct {
	service *car.Service
}

func New(service *car.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidateCarRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateCar(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	resp, err := h.service.GetCar(r.Context(), id)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
	
	var req models.CarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidateCarRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	resp, err := h.service.UpdateCar(r.Context(), id, req)
	if err != nil {
		http.Error(w, "Failed to update car: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	err := h.service.DeleteCar(r.Context(), id)
	if err != nil {

		if strings.Contains(err.Error(), "no car found") {
			 http.Error(w, "Car not found", http.StatusNotFound)
			 return
		}
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}