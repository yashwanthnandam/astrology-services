package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemedyHandler struct {
	usecase usecase.RemedyUsecase
}

func NewRemedyHandler(r *gin.Engine, u usecase.RemedyUsecase) {
	handler := &RemedyHandler{u}
	r.GET("/remedy/rudraksha/:id", handler.GetRudrakshaSuggestion)
	r.GET("/remedy/gemstone/:id", handler.GetGemstoneSuggestion)
}

func (h *RemedyHandler) GetRudrakshaSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	suggestion, err := h.usecase.GetRudrakshaSuggestion(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suggestion)
}

func (h *RemedyHandler) GetGemstoneSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	suggestion, err := h.usecase.GetGemstoneSuggestion(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suggestion)
}
