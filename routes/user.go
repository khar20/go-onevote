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
	auth, ok := sess.Values["authenticated"].(bool)
	id, _ := sess.Values["user-id"].(string)

	if !ok || !auth {
		return c.String(http.StatusUnauthorized, "Please log in to access this page")
	}

	var user *models.User
	for _, u := range models.GetUsers() {
		if u.ID == id {
			user = &u
			break
		}
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
	return sess.Save(c.Request(), c.Response())
}
