package cid /* import "s32x.com/cid/service" */

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// repoTemp is the template that is rendered when a repo is requested
const repoTemp = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.Domain}}/{{.Path}} git {{.UserURL}}/{{.Path}}">
		<meta http-equiv="refresh" content="0; url={{.UserURL}}/{{.Path}}">
	</head>
	<body>{{.Message}}<a href="{{.UserURL}}/{{.Path}}">move along</a>.</body>
</html>`

// renderer is the custom webpage renderer that can be used as an echo.Renderer
var renderer = &Template{template.Must(template.New("repo").Parse(repoTemp))}

// Service contains the required data needed for hosting a custom go import
// domain API
type Service struct{ userURL, domain string }

// New creates and returns a new Service struct reference using the passed
// params
func New(userURL, domain string) *Service {
	return &Service{userURL: userURL, domain: domain}
}

// Echo creates and returns a fully configured echo router that is able to proxy
// git repos for a user
func (s *Service) Echo() *echo.Echo {
	// Initialize the echo Echo and bind middleware
	e := echo.New()
	e.HideBanner = true
	e.Renderer = renderer

	// Bind all middleware
	e.Pre(middleware.RemoveTrailingSlashWithConfig(
		middleware.TrailingSlashConfig{
			RedirectCode: http.StatusPermanentRedirect,
		}))
	e.Pre(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Bind the redirect for the user
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, s.userURL)
	})

	// Bind the redirect for all repositories
	e.GET("/:repo", s.repo)
	e.GET("/:repo/*", s.repo)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	return e
}

// Repo contains the data that will be rendered in the repo redirect template
type Repo struct{ UserURL, Domain, Message, Path string }

// repo handles rendering an html repository redirect page with the proper
// head metadata
func (s *Service) repo(c echo.Context) error {
	return c.Render(http.StatusOK, "repo", &Repo{
		UserURL: s.userURL,
		Domain:  s.domain,
		Message: "Nothing to see here;",
		Path:    c.Param("repo"),
	})
}

// Template is a renderer that contains html templates
type Template struct{ templates *template.Template }

// Render makes the Template type satisfy the echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{},
	c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
