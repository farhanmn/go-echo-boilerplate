package controller

import (
	"go-echo-experiment/internal/model"
	"go-echo-experiment/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"go-echo-experiment/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c echo.Context) error {
	user := &model.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Insert into database
	_, err := h.Service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to register user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully!"})
}

func (h *UserHandler) Login(c echo.Context) error {
	reqUser := &model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if err := c.Bind(reqUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Get user from DB
	dbUser, err := h.Service.LoginUser(reqUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	// Compare Password
	if !utils.Match(reqUser.Password, dbUser.Password, dbUser.Salt) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful!"})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	email := c.Param("email")
	user, err := h.Service.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}
