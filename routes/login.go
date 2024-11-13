package routes

import (
	"net/http"
	"onevote/models"
	"onevote/templates"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// handlers
func GetLoginPage(c echo.Context) error {
	data := templates.LoginData{}

	return Render(c, http.StatusFound, templates.LoginTempl(data))
}

func HandleRedirect(c echo.Context) error {
	data := templates.LoginData{}

	cipParam := c.Param("cip")

	if cipParam == "" {
		return Render(c, http.StatusFound, templates.LoginTempl(data))
	}

	_, err := strconv.Atoi(cipParam)

	if err != nil || len(cipParam) > 6 {
		return c.Redirect(http.StatusFound, "/login")
	}

	if cipParam != "" && len(cipParam) <= 6 {
		data.Cip = cipParam
	}

	return Render(c, http.StatusFound, templates.LoginTempl(data))
}

func PostLogin(c echo.Context) error { // htmx
	cip := c.FormValue("cip")

	if cip == "" {
		return c.HTML(http.StatusBadRequest, "<p style='color: red;'>CIP y contraseña son necesarios</p>")
	}

	user, err := models.GetUserByCIP(cip)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error al recuperar usuario")
	}

	if user == nil {
		return c.String(http.StatusUnauthorized, "CIP o contraseña inválidos")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 1 day
		HttpOnly: true,
	}
	sess.Values["authenticated"] = true
	sess.Values["user-id"] = strconv.Itoa(user.ID)

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	c.Response().Header().Set("HX-Location", "/profile")
	return c.NoContent(http.StatusFound)
}
