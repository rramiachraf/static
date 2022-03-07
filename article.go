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

type Article struct {
	Title   string
	Date    time.Time
	Content []byte
	Slug    string
}

// get articles from articles directory
func getArticles() ([]Article, error) {
	var articles []Article

	if paths, err := filepath.Glob("*.md"); err == nil {
		for _, p := range paths {
			var article Article
			article.parse(p)
			articles = append(articles, article)
		}
	}

	return articles, nil
}

// generate articles to the specified "out" directory
func buildArticles(c Conf, out string) ([]map[string]string, error) {
	articles, err := getArticles()

	if err != nil {
		return nil, err
	}

	for _, a := range articles {
		p := path.Join(out, "article", a.Slug+".html")
		f, err := os.Create(p)

		if err != nil {
			fmt.Println("error creating", p)
			continue
		}

		defer f.Close()

		data := map[string]string{
			"Title":       a.Title,
			"Date":        a.Date.Format("Jan 02 2006"),
			"Description": c.Description,
			"Footer":      c.Footer,
			"Content":     string(a.Content),
		}

		tmpl := parseTheme(c.Theme, "article")
		tmpl.Execute(f, data)
	}

	sort.SliceStable(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})

	var indexArticles []map[string]string

	for _, a := range articles {
		iA := map[string]string{
			"Title": a.Title,
			"Slug":  a.Slug,
			"Date":  a.Date.Format("Jan 02 2006"),
		}
		indexArticles = append(indexArticles, iA)
	}

	return indexArticles, nil
}

// read first two line to parse title and date of the article
func (a *Article) parse(p string) error {
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
			a.Title = sc.Text()
		case 2:
			t, _ := time.Parse("January 02, 2006 - 15:04", sc.Text())
			a.Date = t
		case 3:
			continue
		default:
			markdown = bytes.Join([][]byte{markdown, sc.Bytes()}, []byte("\n"))
		}

	}

	a.Slug = slug.Make(a.Title)
	a.Content = md.Run(markdown)
	return nil
}
