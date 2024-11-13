package routes

import (
	"net/http"

	"onevote/templates"

	"github.com/labstack/echo/v4"
)

// handlers
func GetChatPage(c echo.Context) error {
	return Render(c, http.StatusOK, templates.HomeTempl())
}
