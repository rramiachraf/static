package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/gosimple/slug"
	md "github.com/russross/blackfriday/v2"
)

type post struct {
	title   string
	date    time.Time
	content []byte
	slug    string
}

// get articles from articles directory
func getPosts() ([]post, error) {
	var posts []post

	if paths, err := filepath.Glob("*.md"); err == nil {
		for _, p := range paths {
			var pt post
			pt.parse(p)
			posts = append(posts, pt)
		}
	}

	return posts, nil
}

// generate articles to the specified "out" directory
func buildPosts(c Conf, out string) ([]map[string]string, error) {
	posts, err := getPosts()

	if err != nil {
		return nil, err
	}

	for _, pt := range posts {
		p := path.Join(out, "post", pt.slug+".html")
		f, err := os.Create(p)

		if err != nil {
			fmt.Println("error creating", p)
			continue
		}

		defer f.Close()

		data := map[string]string{
			"Title":       pt.title,
			"Date":        pt.date.Format("Jan 02 2006"),
			"Description": c.Description,
			"Footer":      c.Footer,
			"Content":     string(pt.content),
		}

		tmpl := parseTheme(c.Theme, "post")
		tmpl.Execute(f, data)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].date.After(posts[j].date)
	})

	var nav []map[string]string

	for _, pt := range posts {
		item := map[string]string{
			"Title": pt.title,
			"Slug":  pt.slug,
			"Date":  pt.date.Format("Jan 02 2006"),
		}
		nav = append(nav, item)
	}

	return nav, nil
}

// read first two line to parse title and date of the article
func (pt *post) parse(p string) error {
	f, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(bytes.NewReader(f))
	var markdown []byte

	i := 0
	for sc.Scan() {
		i++
		switch i {
		case 1:
			pt.title = sc.Text()
		case 2:
			t, _ := time.Parse("January 02, 2006 - 15:04", sc.Text())
			pt.date = t
		case 3:
			continue
		default:
			markdown = bytes.Join([][]byte{markdown, sc.Bytes()}, []byte("\n"))
		}

	}

	pt.slug = slug.Make(pt.title)
	pt.content = md.Run(markdown)
	return nil
}
