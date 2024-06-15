package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KundaliHandler struct {
	usecase usecase.KundaliUsecase
}

func NewKundaliHandler(r *gin.Engine, u usecase.KundaliUsecase) {
	handler := &KundaliHandler{u}
	r.GET("/kundali/:id", handler.GetKundali)
}

func (h *KundaliHandler) GetKundali(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	kundali, err := h.usecase.GetKundali(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kundali)
}
