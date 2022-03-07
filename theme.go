package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path"
	"text/template"
)

type Theme map[string][]byte

func openTheme(name string) (Theme, error) {
	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	t := tar.NewReader(f)
	theme := make(Theme)

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
