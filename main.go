package main

import (
	"log"
	"net/http"
	"project-app-inventory-restapi-golang-rahmadhany/database"
	"project-app-inventory-restapi-golang-rahmadhany/handler"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
	"project-app-inventory-restapi-golang-rahmadhany/router"
	"project-app-inventory-restapi-golang-rahmadhany/service"
)

func main() {
	conn, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	repo := repository.NewRepository(conn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	r := router.NewRouter(h)

	log.Println("server starting on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
