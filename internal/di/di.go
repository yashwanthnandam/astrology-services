package di

import (
	"astrology-services/config"
	mysqlRepository "astrology-services/internal/repository/mysql"
	"astrology-services/internal/services"
	"astrology-services/internal/usecase"
	"astrology-services/pkg/gpt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Container struct {
	KundaliUsecase     usecase.KundaliUsecase
	AshtakvargaUsecase usecase.AshtakvargaUsecase
	DashaUsecase       usecase.DashaUsecase
	ReportUsecase      usecase.ReportUsecase
	RemedyUsecase      usecase.RemedyUsecase
	DoshaUsecase       usecase.DoshaUsecase
	MatchUsecase       usecase.MatchUsecase
	GPTService         *gpt.GPTService
}

func NewContainer(cfg *config.Config) *Container {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	kundaliRepo := mysqlRepository.NewKundaliRepository(db)
	ashtakvargaRepo := mysqlRepository.NewAshtakvargaRepository(db)
	dashaRepo := mysqlRepository.NewDashaRepository(db)
	reportRepo := mysqlRepository.NewReportRepository(db)
	remedyRepo := mysqlRepository.NewRemedyRepository(db)
	doshaRepo := mysqlRepository.NewDoshaRepository(db)
	matchRepo := mysqlRepository.NewMatchRepository(db)

	kundaliUsecase := usecase.NewKundaliUsecase(kundaliRepo)
	ashtakvargaUsecase := usecase.NewAshtakvargaUsecase(ashtakvargaRepo)
	dashaUsecase := usecase.NewDashaUsecase(dashaRepo)
	reportService := services.NewAstrologyReportService()
	reportUsecase := usecase.NewReportUsecase(reportRepo, reportService)
	remedyUsecase := usecase.NewRemedyUsecase(remedyRepo)
	doshaUsecase := usecase.NewDoshaUsecase(doshaRepo)
	matchUsecase := usecase.NewMatchUsecase(matchRepo)

	gptService := gpt.NewGPTService(cfg.GPTAPIKey)

	return &Container{
		KundaliUsecase:     *kundaliUsecase,
		AshtakvargaUsecase: *ashtakvargaUsecase,
		DashaUsecase:       *dashaUsecase,
		ReportUsecase:      *reportUsecase,
		RemedyUsecase:      *remedyUsecase,
		DoshaUsecase:       *doshaUsecase,
		MatchUsecase:       *matchUsecase,
		GPTService:         gptService,
	}
}
