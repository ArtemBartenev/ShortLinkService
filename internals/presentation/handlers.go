package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shortlink/internals/domain/entity"
	"shortlink/internals/domain/service"
)

type Request struct {
	OriginalURL string `json:"originalURL"`
}

type Response struct {
	ShortURL string `json:"shortURL"`
}

type Handler struct {
	service *service.LinkService
}

func NewHandler(service *service.LinkService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.createShortURL(w, r)
	} else if r.Method == "GET" {
		h.getOriginalURL(w, r)
	}
}

func (h *Handler) createShortURL(w http.ResponseWriter, r *http.Request) {
	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if request.OriginalURL == "" {
		http.Error(w, "Field \"OriginalURL\" cant be empty.", http.StatusUnprocessableEntity)
	}

	link := entity.NewLink()
	link.OriginalURL = request.OriginalURL
	link.ShortURL = fmt.Sprintf("http://%s/", r.Host)

	ShortURL, err := h.service.ConvertOriginalURLToShortURL(link)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Response{ShortURL: ShortURL})
}

func (h *Handler) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	ShortURL := r.URL.Query().Get("shortURL")
	if ShortURL == "" {
		http.Error(w, "ShortURL parameter is missing", http.StatusBadRequest)
	}

	link := entity.NewLink()
	link.ShortURL = ShortURL
	OriginalURL, err := h.service.GetOriginalURLByShortURL(link)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if OriginalURL == "" {
		http.Error(w, "Original URL not found.", http.StatusNotFound)
	}

	http.Redirect(w, r, OriginalURL, http.StatusMovedPermanently)
}
