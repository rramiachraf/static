package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gosimple/slug"
	"github.com/russross/blackfriday/v2"
)

const timeLayout = "02/01/2006 15:04"

type post struct {
	Title   string
	Date    time.Time
	Slug    string
	Content []byte
}

func newPost() post {
	var p post
	return p
}

func (post *post) parse(p string) error {
	file, err := os.Open(p)
	if err != nil {
		return err
	}

	defer file.Close()

	metadata, content := scan(file)
	m := parseMetadata(metadata)

	post.Title = m["TITLE"]
	post.Slug = slug.Make(m["TITLE"])
	post.Content = blackfriday.Run(highlightCode(content))
	if date, err := time.Parse(timeLayout, m["DATE"]); err == nil {
		post.Date = date
	}

	return nil
}

// scan file and seperate metadata from content
func scan(f io.Reader) ([]byte, []byte) {
	var metadata []byte
	var content []byte

	s := bufio.NewScanner(f)
	var found bool

	for s.Scan() {
		if bytes.Compare(s.Bytes(), []byte{}) == 0 {
			found = true
		}

		switch found {
		case true:
			content = bytes.Join([][]byte{content, s.Bytes()}, []byte("\n"))
		case false:
			metadata = bytes.Join([][]byte{metadata, s.Bytes()}, []byte("\n"))
		}
	}

	return metadata, content
}

// parse metadata from bytes into a string map
func parseMetadata(metadata []byte) map[string]string {
	title := regexp.MustCompile(`TITLE\s(.*)`).FindSubmatch(metadata)
	date := regexp.MustCompile(`DATE\s(.*)`).FindSubmatch(metadata)

	return map[string]string{
		"TITLE": string(title[1]),
		"DATE":  string(date[1]),
	}
}

// search the current directory for posts (files ending with .md)
func getPosts() ([]string, error) {
	return filepath.Glob("*.md")
}
