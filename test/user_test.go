package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-echo-experiment/internal/controller"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/test@example.com", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := controller.UserHandler{} // Simulasikan controller
	err := h.GetUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
