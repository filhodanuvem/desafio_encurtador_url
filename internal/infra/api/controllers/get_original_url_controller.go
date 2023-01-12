package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/desafio_encurtador_url/internal/usecases"
)

type GetOriginalURLController struct {
	GetOriginalURL *usecases.GetOriginalURL
}

func NewGetOriginalURLController(getOriginalURL *usecases.GetOriginalURL) *GetOriginalURLController {
	return &GetOriginalURLController{
		GetOriginalURL: getOriginalURL,
	}
}

func (m *GetOriginalURLController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	shortURL := r.URL.Query().Get("code")
	output, err := m.GetOriginalURL.Execute(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(output)
}
