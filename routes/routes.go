package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := component.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

//func createSession(c echo.Context) error {
//	sess, err := session.Get("session", c)
//	if err != nil {
//		return err
//	}
//
//	sess.Options = &sessions.Options{
//		Path:     "/",
//		MaxAge:   86400,
//		HttpOnly: true,
//	}
//
//	sess.Values["authenticated"] = true
//	return sess.Save(c.Request(), c.Response())
//}

//func clearSession(c echo.Context) error {
//	sess, err := session.Get("session", c)
//	if err != nil {
//		return err
//	}
//
//	sess.Options.MaxAge = -1
//	return sess.Save(c.Request(), c.Response())
//}

// middleware user auth check session
func voterCheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		auth, ok := sess.Values["authenticated"].(bool)
		if !ok || !auth {
			return c.String(http.StatusUnauthorized, "Please log in to access this page")
		}
		return next(c)
	}
}

// middleware user auth check session
func adminCheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		auth, ok := sess.Values["authenticated"].(bool)
		if !ok || !auth {
			return c.String(http.StatusUnauthorized, "Please log in to access this page")
		}
		return next(c)
	}
}

func SetUpRoutes(e *echo.Echo) {
	e.GET("/", GetHomePage)                                 // render homepage
	e.GET("/login/:cip", GetLoginPage)                      // render login page with CIP
	e.GET("/login", GetLoginPage)                           // render login page
	e.GET("/profile", GetUserProfile, voterCheckAuth)       // render user profile
	e.GET("/candidates", GetCandidatesPage)                 // render candidates list
	e.GET("/candidates/:candidate-id", GetCandidateProfile) // render specific candidate profile
	e.GET("/vote", GetVotePage, voterCheckAuth)             // render voting page
	e.GET("/admin", GetAdminPage, adminCheckAuth)           // render blockchain or chain info
	e.GET("/chat", GetChatPage, voterCheckAuth)             // render chatbot page
	e.GET("/chain", GetChainPage, voterCheckAuth)           // render blockchain or chain info

	e.POST("/verify-cip", PostVerifyCIP)          // htmx - verify CIP
	e.POST("/logout", PostLogout, voterCheckAuth) // htmx - handle session logout
	e.POST("/login", PostLogin)                   // htmx - handle login
	e.POST("/vote", PostVote, voterCheckAuth)     // htmx - handle post vote
}
