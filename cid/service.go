package cid /* import "s32x.com/cid/cid" */

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start starts the import proxy service using the passed configuration vars
func Start(redirect, userURL, domain, port string) {
	// Initialize the echo Echo and bind middleware
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Configure the renderer for static pages and hide the echo banner
	e.HideBanner = true
	e.Renderer = renderer

	// Configure the domain redirect on the index
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, redirect)
	})

	// Bind the redirect for all repositories
	e.GET("/:repo", func(c echo.Context) error {
		return c.Render(http.StatusOK, "", &Data{
			UserURL: userURL,
			Domain:  domain,
			Message: "Nothing to see here;",
			Path:    c.Param("repo"),
		})
	})
	e.GET("/:repo/*", func(c echo.Context) error {
		return c.Render(http.StatusOK, "", &Data{
			UserURL: userURL,
			Domain:  domain,
			Message: "Nothing to see here;",
			Path:    c.Param("repo"),
		})
	})
	e.Logger.Fatal(e.Start(":" + port))
}
