package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo *ProductRepo
}

func NewHandler(repo *ProductRepo) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Repo.List())
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := h.Repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	created := h.Repo.Create(p)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	updated, err := h.Repo.Update(id, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
