package routes

import (
	"net/http"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/controllers"
)

type ShortenerRoute struct {
	Mux                      *http.ServeMux
	MakeShortenerController  *controllers.MakeShortenerURLController
	GetOriginalURLController *controllers.GetOriginalURLController
}

func NewShortenerRoute(mux *http.ServeMux, makeShortenerController *controllers.MakeShortenerURLController, getOriginalURLController *controllers.GetOriginalURLController) *ShortenerRoute {
	return &ShortenerRoute{
		Mux:                      mux,
		MakeShortenerController:  makeShortenerController,
		GetOriginalURLController: getOriginalURLController,
	}
}

func (s *ShortenerRoute) Init() {
	s.Mux.HandleFunc("/cut", s.MakeShortenerController.Handle)
	s.Mux.HandleFunc("/uncut", s.GetOriginalURLController.Handle)
}
