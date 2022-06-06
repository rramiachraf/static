package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Title       string
	Description string
	URL         string
	Theme       string
}

func parseConfig(p string) (conf, error) {
	var c conf

	f, err := os.Open(p)
	if err != nil {
		return c, err
	}

	defer f.Close()

	d := yaml.NewDecoder(f)
	d.Decode(&c)

	return c, nil
}
