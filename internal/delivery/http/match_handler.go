package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	usecase usecase.MatchUsecase
}

func NewMatchHandler(r *gin.Engine, u usecase.MatchUsecase) {
	handler := &MatchHandler{u}
	r.GET("/match/:id", handler.GetKundaliMatching)
}

func (h *MatchHandler) GetKundaliMatching(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	matching, err := h.usecase.GetKundaliMatching(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matching)
}
