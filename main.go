package main

import (
	"flag"
	"log"
)

func main() {
	config := flag.String("config", "config.yml", "config file path")
	out := flag.String("out", "dist", "directory path where the generated files will be saved")

	flag.Parse()

	c, err := parseConfig(*config)
	if err != nil {
		log.Fatalln(err)
	}

	build(c, *out)
}
