package blogrenderer

import (
	"io"
	"text/template"

	blogposts "github.com/tmanley1985/learn-go-with-tests/reading-files"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.New("blog").Parse(postTemplate)

	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
