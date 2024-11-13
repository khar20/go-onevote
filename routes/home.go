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
	formCIP := c.FormValue("cip")

	user, err := models.GetUserByCIP(formCIP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch user"})
	}

	if user == nil {
		return c.HTML(http.StatusBadRequest, "<p style='color: red;'>CIP no válido</p>")
	}

	c.Response().Header().Set("HX-Location", fmt.Sprintf("/login/%s", user.CIP))
	return c.NoContent(http.StatusFound)
}
