package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Title       string
	Description string
	URL         string
	Theme       string
}

func Parse(p string) (Conf, error) {
	var c Conf

	f, err := os.Open(p)
	if err != nil {
		return c, err
	}

	defer f.Close()

	d := yaml.NewDecoder(f)
	d.Decode(&c)

	return c, nil
}
