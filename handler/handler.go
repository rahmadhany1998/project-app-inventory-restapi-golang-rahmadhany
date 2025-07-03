package handler

import "project-app-inventory-restapi-golang-rahmadhany/service"

type Handler struct {
	UserHandler UserHandler
}

func NewHandler(s service.Service) Handler {
	return Handler{
		UserHandler: NewUserHandler(s),
	}
}
