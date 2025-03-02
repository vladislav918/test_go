package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	repository "test-project/internal/database/postgres"
	"test-project/internal/models"
)

type MeasureHandler struct {
	repo repository.Measure
}

func NewMeasureHandler(repo repository.Measure) *MeasureHandler {
	return &MeasureHandler{repo: repo}
}

func (h *MeasureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/measure/")

	switch r.Method {
	case http.MethodGet:
		if path == "" {
			h.GetAllMeasures(w, r)
		} else {
			h.GetMeasureByID(w, r, path)
		}
	case http.MethodPost:
		h.CreateMeasure(w, r)
	case http.MethodPut:
		h.UpdateMeasure(w, r, path)
	case http.MethodDelete:
		h.DeleteMeasure(w, r, path)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MeasureHandler) GetAllMeasures(w http.ResponseWriter, r *http.Request) {
	measures, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(measures)
}

func (h *MeasureHandler) GetMeasureByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid measure ID", http.StatusBadRequest)
		return
	}

	measure, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, "Measure not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(measure)
}

func (h *MeasureHandler) CreateMeasure(w http.ResponseWriter, r *http.Request) {
	var measure models.Measure
	if err := json.NewDecoder(r.Body).Decode(&measure); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&measure); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"id": measure.ID})
}

func (h *MeasureHandler) UpdateMeasure(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid measure ID", http.StatusBadRequest)
		return
	}

	var measure models.Measure
	if err := json.NewDecoder(r.Body).Decode(&measure); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	measure.ID = id
	if err := h.repo.Update(&measure); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(measure)
}

func (h *MeasureHandler) DeleteMeasure(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid measure ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
