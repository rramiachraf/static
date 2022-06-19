package main

import (
	"flag"
	"log"

	"github.com/rramiachraf/static/config"
)

func main() {
	configPath := flag.String("config", "config.yml", "config file path")
	out := flag.String("out", "dist", "directory path where the generated files will be saved")

	flag.Parse()

	c, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	build(c, *out)
}
