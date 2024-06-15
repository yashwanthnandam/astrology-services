package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DoshaHandler struct {
	usecase usecase.DoshaUsecase
}

func NewDoshaHandler(r *gin.Engine, u usecase.DoshaUsecase) {
	handler := &DoshaHandler{u}
	r.GET("/dosha/:id", handler.GetDoshaAnalysis)
}

func (h *DoshaHandler) GetDoshaAnalysis(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	analysis, err := h.usecase.GetDoshaAnalysis(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, analysis)
}
