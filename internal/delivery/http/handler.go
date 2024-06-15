package http

import (
	"astrology-services/internal/di"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, container *di.Container) {
	NewKundaliHandler(r, container.KundaliUsecase)
	NewAshtakvargaHandler(r, container.AshtakvargaUsecase)
	NewDashaHandler(r, container.DashaUsecase)
	NewReportHandler(r, container.ReportUsecase)
	NewRemedyHandler(r, container.RemedyUsecase)
	NewDoshaHandler(r, container.DoshaUsecase)
	NewMatchHandler(r, container.MatchUsecase)
}
