package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ListProducts())
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	created := CreateProduct(p)
	json.NewEncoder(w).Encode(created)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	updated, err := UpdateProduct(id, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
