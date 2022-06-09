package main

import (
	"embed"
	"os"
	"path"
	"text/template"
)

//go:embed classic
var defaultTheme embed.FS

type theme map[string][]byte

func (t theme) openDefaultTheme(dist string) error {
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

		if f.IsDir() {
			dirname := path.Join("classic", f.Name())
			os.Mkdir(path.Join(dist, f.Name()), 0777)
			copyDir(dirname, defaultTheme, dist)
		}
	}

	return nil
}

func copyDir(dirname string, src embed.FS, dist string) error {
	dir, err := src.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, file := range dir {
		data, err := src.ReadFile(path.Join(dirname, file.Name()))
		if err != nil {
			continue
		}

		err = os.WriteFile(path.Join(dist, path.Base(dirname), file.Name()), data, 0777)
		if err != nil {
			continue
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
