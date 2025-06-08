package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santigorbe/hexagonal_arq/internal/adapter/inbound/http"
	"github.com/santigorbe/hexagonal_arq/internal/adapter/outbound/memory"
	"github.com/santigorbe/hexagonal_arq/internal/usecase"
)

func main() {
	r := gin.Default()

	repo := memory.NewInMemoryUserRepo()
	service := usecase.NewUserService(repo)
	handler := http.NewUserHandler(service)

	handler.RegisterRoutes(r)
	r.Run(":8080")
}

//
