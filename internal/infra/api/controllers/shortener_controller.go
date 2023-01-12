package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/desafio_encurtador_url/internal/usecases"
)

type MakeShortenerURLController struct {
	MakeShortenerURL *usecases.MakeShortenerURL
}

func NewShortenerController(makeShortenerURL *usecases.MakeShortenerURL) *MakeShortenerURLController {
	return &MakeShortenerURLController{
		MakeShortenerURL: makeShortenerURL,
	}
}

func (m *MakeShortenerURLController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	var input usecases.MakeShortenerInputURL
	json.NewDecoder(r.Body).Decode(&input)
	output, err := m.MakeShortenerURL.Execute(input)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
