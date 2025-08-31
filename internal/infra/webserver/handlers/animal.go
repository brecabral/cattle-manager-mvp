package handlers

import "net/http"

type AnimalHandler struct{}

func (h *AnimalHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AnimalHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AnimalHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AnimalHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AnimalHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
