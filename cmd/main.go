package main

import (
	"astrology-services/config"
	httpDelivery "astrology-services/internal/delivery/http"
	"astrology-services/internal/di"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	container := di.NewContainer(cfg)
	router := gin.Default()

	httpDelivery.NewRouter(router, container)

	log.Fatal(router.Run(cfg.ServerAddress))
}
