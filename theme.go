package main

import (
	"embed"
	"path"
	"text/template"
)

//go:embed classic
var defaultTheme embed.FS

type theme map[string][]byte

func (t theme) openDefaultTheme() error {
	d, err := defaultTheme.ReadDir("classic")
	if err != nil {
		return nil
	}

	for _, f := range d {
		if !f.IsDir() {
			name := f.Name()
			r, _ := defaultTheme.ReadFile(path.Join("classic", name))
			t[name] = r
		}
	}

	return nil
}

func (t theme) parse(theme, page string) *template.Template {
	tmpl := template.New(page)
	tmpl.Parse(string(t[page+".tmpl"]))
	tmpl.Parse(string(t["head.tmpl"]))
	tmpl.Parse(string(t["footer.tmpl"]))

	return tmpl
}
