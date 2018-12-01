package cid /* import "s32x.com/cid" */

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// renderer is the custom webpage renderer that can be used as an echo.Renderer
var renderer = &Template{template.Must(template.New("").Parse(temp))}

// Data contains the data that will be rendered in the template
type Data struct{ UserURL, Domain, Message, Path string }

// temp is the template that is rendered when a repo is requested
const temp = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.Domain}}{{.Path}} git https://{{.UserURL}}{{.Path}}">
		<meta http-equiv="refresh" content="0; url=http://{{.Domain}}">
	</head>
	<body>{{.Message}}<a href="http://{{.Domain}}">move along</a>.</body>
</html>`

// Template is a renderer that contains html templates
type Template struct{ templates *template.Template }

// Render makes the Template type satisfy the echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{},
	c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
