package main

import (
	"log"
	"net/http"
	"project-app-inventory-restapi-golang-rahmadhany/database"
	"project-app-inventory-restapi-golang-rahmadhany/handler"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
	"project-app-inventory-restapi-golang-rahmadhany/router"
	"project-app-inventory-restapi-golang-rahmadhany/service"
	"project-app-inventory-restapi-golang-rahmadhany/utils"

	"go.uber.org/zap"
)

func main() {

	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal("cant read file .env")
	}

	// init logger
	logger, err := utils.InitLogger("./logs/app-", config)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	conn, err := database.InitDB(config)
	if err != nil {
		logger.Fatal("Database error: %v", zap.Error(err))
	}

	repo := repository.NewRepository(conn, logger)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, config)

	r := router.NewRouter(h)

	log.Println("server starting on port : 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal("can't run service", zap.Error(err))
	}
}
