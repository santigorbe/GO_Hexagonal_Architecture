package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santigorbe/hexagonal_arq/internal/domain"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/users", h.GetAll)
	r.POST("/users", h.Create)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users := h.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.CreateUser(user)
	c.JSON(http.StatusCreated, created)
}
