package post

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
func ParseHeader(header []byte) map[string]string {
	title := regexp.MustCompile(`TITLE\s(.*)`).FindSubmatch(header)
	date := regexp.MustCompile(`DATE\s(.*)`).FindSubmatch(header)

	return map[string]string{
		"TITLE": string(title[1]),
		"DATE":  string(date[1]),
	}
}

// search the current directory for posts (files ending with .md)
func GetAll(dir string) ([]string, error) {
	glob := "*.md"
	if dir != "" {
		glob = path.Join(dir, "*.md")
	}

	return filepath.Glob(glob)
}
