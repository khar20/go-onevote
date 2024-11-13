package routes

import (
	"net/http"
	"onevote/models"
	"onevote/templates"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// handlers
func GetUserProfile(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sessAuth, okAuth := sess.Values["authenticated"].(bool)
	sessID, _ := sess.Values["user-id"].(string)

	if !okAuth || !sessAuth {
		return c.String(http.StatusUnauthorized, "Please log in to access this page")
	}

	user, err := models.GetUserByID(sessID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving user profile")
	}

	if user == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	data := templates.UserProfileData{
		User: user,
	}

	return Render(c, http.StatusOK, templates.UserProfileTempl(data))
}

func PostLogout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["authenticated"] = false
	sess.Values["user-id"] = "0"
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	c.Response().Header().Set("HX-Location", "/login")
	return c.NoContent(http.StatusFound)
}
