package api

import (
	"io"
	"net/http"

	"github.com/AnotherOneDushnila/KV_Storage/internal/store"
)



type Handler struct {
	store store.Store
}



func NewHandler(s store.Store) *Handler {
	return &Handler{store: s}
}


func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}


func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	collection := r.URL.Query().Get("collection")
	key := r.URL.Query().Get("key")

	if collection == "" || key == "" {
		http.Error(w, "collection and key required!", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := h.store.Put(collection, key, body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}


func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	collection := r.URL.Query().Get("collection")
	key := r.URL.Query().Get("key")

	if collection == "" || key == "" {
		http.Error(w, "collection and key required!", http.StatusBadRequest)
		return
	}

	val, err := h.store.Get(collection, key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(val)
}


func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	collection := r.URL.Query().Get("collection")
	key := r.URL.Query().Get("key")

	if collection == "" || key == "" {
		http.Error(w, "collection and key required!", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(collection, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}