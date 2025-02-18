package controller

import (
	"go-echo-experiment/internal/model"
	"go-echo-experiment/pkg/response"
	"go-echo-experiment/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"go-echo-experiment/internal/service"
)

var (
	successJSON = response.Success
	errorJSON   = response.Error
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
		return errorJSON(c, http.StatusBadRequest, "Invalid input")
	}

	// Insert into database
	_, err := h.Service.CreateUser(user)
	if err != nil {
		return errorJSON(c, http.StatusInternalServerError, "Failed to register user")
	}

	return successJSON(c, http.StatusCreated, "User registered successfully!", nil)
}

func (h *UserHandler) Login(c echo.Context) error {
	reqUser := &model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if err := c.Bind(reqUser); err != nil {
		return errorJSON(c, http.StatusBadRequest, "Invalid input")
	}

	// Get user from DB
	dbUser, err := h.Service.LoginUser(reqUser)
	if err != nil {
		return errorJSON(c, http.StatusUnauthorized, "Invalid email or password")
	}

	// Compare Password
	if !utils.Match(reqUser.Password, dbUser.Password, dbUser.Salt) {
		return errorJSON(c, http.StatusUnauthorized, "Invalid email or password")
	}

	return successJSON(c, http.StatusOK, "Login successful!", nil)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	email := c.Param("email")
	user, err := h.Service.GetUserByEmail(email)
	if err != nil {
		return errorJSON(c, http.StatusNotFound, "User not found")
	}
	return successJSON(c, http.StatusOK, "", user)
}
