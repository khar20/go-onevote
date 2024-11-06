package routes

import (
	"fmt"
	"net/http"

	"onevote/models"
	"onevote/templates"

	"github.com/labstack/echo/v4"
)

// handlers
func GetHomePage(c echo.Context) error {
	return Render(c, http.StatusOK, templates.HomeTempl())
}

func PostVerifyCIP(c echo.Context) error { //htmx
	formCip := c.FormValue("cip")

	for _, user := range models.GetUsers() {
		if user.Cip == formCip {
			c.Response().Header().Set("HX-Location", fmt.Sprintf("/login/%s", user.Cip))
			return c.NoContent(http.StatusFound)
		}
	}

	return c.HTML(http.StatusBadRequest, "<p style='color: red;'>CIP no válido</p>")
}
