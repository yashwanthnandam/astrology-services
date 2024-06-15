package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashaHandler struct {
	usecase usecase.DashaUsecase
}

func NewDashaHandler(r *gin.Engine, u usecase.DashaUsecase) {
	handler := &DashaHandler{u}
	r.GET("/dasha/:id", handler.GetDasha)
}

func (h *DashaHandler) GetDasha(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	dasha, err := h.usecase.GetDasha(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dasha)
}
