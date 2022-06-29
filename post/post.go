package post

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gosimple/slug"
	"github.com/rramiachraf/static/highlight"
	"github.com/russross/blackfriday/v2"
)

const timeLayout = "02/01/2006 15:04"

type Post struct {
	Title   string
	Date    time.Time
	Slug    string
	Content []byte
}

func New() Post {
	var p Post
	return p
}

func (post *Post) Parse(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}

	defer f.Close()

	h, c, err := Scan(f)
	if err != nil {
		return err
	}

	m := ParseHeader(h)

	post.Title = m["TITLE"]
	post.Slug = slug.Make(m["TITLE"])
	post.Content = blackfriday.Run(highlight.Highlight(c))
	if date, err := time.Parse(timeLayout, m["DATE"]); err == nil {
		post.Date = date
	}

	return nil
}

// scan file and seperate header from content
func Scan(r io.Reader) ([]byte, []byte, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}

	header, content, found := bytes.Cut(data, []byte("\n\n"))
	if !found {
		return nil, nil, errors.New("can't parse headers and content from file")
	}

	return header, content, nil
}

// parse header bytes into a map
func ParseHeader(h []byte) map[string]string {
	header := make(map[string]string)
	b := bufio.NewScanner(bytes.NewReader(h))

	for b.Scan() {
		name, value, found := bytes.Cut(b.Bytes(), []byte(" "))
		if found {
			header[string(name)] = string(value)
		}
	}

	return header
}

// search the current directory for posts (files ending with .md)
func GetAll(dir string) ([]string, error) {
	glob := "*.md"
	if dir != "" {
		glob = path.Join(dir, "*.md")
	}

	return filepath.Glob(glob)
}
