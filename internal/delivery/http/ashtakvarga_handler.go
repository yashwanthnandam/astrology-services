package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AshtakvargaHandler struct {
	usecase usecase.AshtakvargaUsecase
}

func NewAshtakvargaHandler(r *gin.Engine, u usecase.AshtakvargaUsecase) {
	handler := &AshtakvargaHandler{u}
	r.GET("/ashtakvarga/:id", handler.GetAshtakvarga)
}

func (h *AshtakvargaHandler) GetAshtakvarga(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ashtakvarga, err := h.usecase.GetAshtakvarga(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ashtakvarga)
}
