package routes

import (
	"net/http"
	"onevote/models"
	"onevote/templates"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetAdminPage(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sessAuth, okAuth := sess.Values["authenticated"].(bool)
	sessAdmin, okAdmin := sess.Values["admin"].(bool)
	sessID, okID := sess.Values["user-id"].(string)

	if !sessAuth || !sessAdmin || !okAuth || !okAdmin || !okID {
		return c.Redirect(http.StatusFound, "/login")
	}

	user, err := models.GetUserByID(sessID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving user profile")
	}

	if user == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	if user.Role != "ADMIN" {
		return c.String(http.StatusForbidden, "Acceso denegado: Solo el administrador puede acceder a esta página.")
	}

	data := templates.TimerData{
		FechaInicio: "",
		FechaFinal:  "",
	}

	return Render(c, http.StatusOK, templates.TimerTempl(data))
}
