package http

import (
	"astrology-services/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	usecase usecase.ReportUsecase
}

func NewReportHandler(r *gin.Engine, u usecase.ReportUsecase) {
	handler := &ReportHandler{u}
	r.GET("/report/general/:id", handler.GetGeneralReport)
	r.GET("/report/planetary", handler.GetPlanetaryReport)
	r.GET("/report/vimshottari", handler.GetVimshottariReport)
	r.GET("/report/yoga", handler.GetYogaReport)
}

func (h *ReportHandler) GetGeneralReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	report, err := h.usecase.GetGeneralReport(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GetPlanetaryReport(c *gin.Context) {
	name := c.Query("name")
	dob := c.Query("dob")
	tob := c.Query("tob")
	placeOfBirth := c.Query("place_of_birth")

	if name == "" || dob == "" || tob == "" || placeOfBirth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	report, err := h.usecase.GeneratePlanetaryReport(name, dob, tob, placeOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GetVimshottariReport(c *gin.Context) {
	name := c.Query("name")
	dob := c.Query("dob")
	tob := c.Query("tob")
	placeOfBirth := c.Query("place_of_birth")

	if name == "" || dob == "" || tob == "" || placeOfBirth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	report, err := h.usecase.GenerateVimshottariReport(name, dob, tob, placeOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GetYogaReport(c *gin.Context) {
	name := c.Query("name")
	dob := c.Query("dob")
	tob := c.Query("tob")
	placeOfBirth := c.Query("place_of_birth")

	if name == "" || dob == "" || tob == "" || placeOfBirth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	report, err := h.usecase.GenerateYogaReport(name, dob, tob, placeOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
