package pkg

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.srv.GetAllUsers()
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve users")
	}

	return c.JSON(http.StatusOK, users)
}
func (h *Handler) Createuser(c echo.Context) error {
	var user Database

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	CreatedUser, err := h.srv.CreateUserdata(user.Name, user.Email, user.Gender)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, CreatedUser)
}
