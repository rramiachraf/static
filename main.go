package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Footer      string `yaml:"footer"`
	Theme       string `yaml:"theme"`
}

func main() {
	var config string
	var out string

	buildSet := flag.NewFlagSet("build", flag.ExitOnError)
	buildSet.StringVar(&config, "config", "config.yml", "config file path")
	buildSet.StringVar(&out, "out", "dist", "directory path where the generated files will be saved")

	if len(os.Args) < 2 {
		fmt.Println("a subcommand must be provided")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		buildSet.Parse(os.Args[2:])
		c := readConf(config)
		build(c, out)
	}
}

func readConf(p string) Conf {
	f, err := os.Open(p)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	var conf Conf
	y := yaml.NewDecoder(f)
	y.Decode(&conf)

	return conf
}

func build(c Conf, out string) {
	theme, err := openTheme(c.Theme)

	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(path.Join(out, "post"), 0777)

	if errors.Is(err, os.ErrExist) {
		fmt.Printf("\"%s\" directory already exist, please remove or specify another one\n", out)
		os.Exit(1)
	}

	posts, err := buildPosts(c, out)
	if err != nil {
		log.Fatalln(err)
	}
	buildIndex(c, posts, out)
	buildStyle(theme["style.css"], out)
}

type index struct {
	Title       string
	Description string
	Footer      string
	Posts       interface{}
}

// build index.html file to "out" directory
func buildIndex(c Conf, posts []map[string]string, out string) error {
	p := path.Join(out, "index.html")
	f, err := os.Create(p)

	if err != nil {
		return err
	}

	defer f.Close()

	var data index
	data.Title = c.Title
	data.Description = c.Description
	data.Footer = c.Footer
	data.Posts = posts

	tmpl := parseTheme(c.Theme, "index")
	tmpl.Execute(f, data)
	return nil
}

// build style.css file to "out" directory
func buildStyle(theme []byte, out string) {
	os.WriteFile(path.Join(out, "style.css"), theme, 0777)
}
