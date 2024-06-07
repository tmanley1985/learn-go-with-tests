package blogrenderer

import (
	"fmt"
	"io"

	blogposts "github.com/tmanley1985/learn-go-with-tests/reading-files"
)

func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "<ul>")

	if err != nil {
		return err
	}

	for _, tag := range p.Tags {
		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)

		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "</ul>")

	if err != nil {
		return err
	}

	return nil
}
