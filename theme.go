package main

import (
	"archive/tar"
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"text/template"
)

//go:embed classic
var defaultTheme embed.FS

type Theme map[string][]byte

func openTheme(name string) (Theme, error) {
	theme := make(Theme)

	if name == "" {
		files := []string{"index.tmpl", "head.tmpl", "footer.tmpl", "post.tmpl", "style.css"}

		for _, f := range files {
			if b, err := defaultTheme.ReadFile("classic/" + f); err == nil {
				theme[f] = b
			}
		}

		return theme, nil
	}

	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	t := tar.NewReader(f)

	for {
		h, err := t.Next()

		if err == io.EOF {
			break
		}

		theme.passFile(path.Base(h.Name), t)
	}

	return theme, nil
}

func (theme Theme) passFile(filename string, reader io.Reader) {
	b, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	theme[filename] = b
}

func parseTheme(theme, page string) *template.Template {
	t, err := openTheme(theme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tmpl := template.New(page)
	tmpl.Parse(string(t[page+".tmpl"]))
	tmpl.Parse(string(t["head.tmpl"]))
	tmpl.Parse(string(t["footer.tmpl"]))

	return tmpl
}
