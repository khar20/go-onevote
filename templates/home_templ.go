// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HomeTempl() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"es\"><head><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,400;0,600;1,100&amp;display=swap\" rel=\"stylesheet\"><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>OneVote - Inicio</title><link rel=\"stylesheet\" href=\"/static/styles.css\"><script src=\"https://unpkg.com/htmx.org@2.0.3\"></script></head><body><nav class=\"navbar\"><div class=\"navbar-logo\"><a href=\"#\">OneVote</a></div><div class=\"navbar-links\"><a href=\"/candidates\">Candidatos</a></div></nav><div class=\"container\"><div class=\"hero-section\"><div class=\"hero-content\"><h1>Bienvenido a OneVote</h1><p>Plataforma de voto no presencial basada en blockchain</p></div></div><div class=\"login-section\"><div class=\"login-form\"><h2>Login</h2><form action=\"#\" hx-post=\"/verify-cip\" hx-target=\"#login-error\"><div class=\"form-group\"><label for=\"CIP\">CIP:</label> <input type=\"text\" id=\"cip\" name=\"cip\" placeholder=\"Ingresa tu CIP\" required></div><button type=\"submit\" class=\"btn\">Login</button></form><div id=\"login-error\"></div></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
