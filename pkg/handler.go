package pkg

import (
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
	users := h.srv.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}
func (h *Handler) Createuser(c echo.Context) error {
	var user Database

	err := c.Bind(&user)
	if err != nil {
		return err
	}
	CreatedUser, err := h.srv.CreateUserdata(user.Name, user.Email, user.Gender)
	return c.JSON(http.StatusCreated, CreatedUser)
}
