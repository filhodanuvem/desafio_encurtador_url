package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (m *MakeShortenerURLController) Handle(c *gin.Context) {
	var input usecases.MakeShortenerInputURL
	json.NewDecoder(c.Request.Body).Decode(&input)
	output, err := m.MakeShortenerURL.Execute(input)
	if err == nil {
		c.JSON(http.StatusCreated, output)
		return
	}
	if err.Error() == "url is required" ||
		err.Error() == "expired url" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"status":  http.StatusInternalServerError,
	})
}
