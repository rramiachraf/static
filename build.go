package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/rramiachraf/static/config"
	"github.com/rramiachraf/static/post"
	"github.com/rramiachraf/static/rss"
)

const perm = 0444

type indexParams struct {
	Title       string
	Description string
	Posts       []post.Post
}

type postParams struct {
	Title       string
	Date        time.Time
	Content     string
	Description string
}

func build(c config.Conf, dist string) {
	os.RemoveAll(dist)
	if err := os.MkdirAll(path.Join(dist, "post"), 0777); err != nil {
		log.Fatalln(err)
	}

	mds, err := post.GetAll("")
	if err != nil {
		log.Fatalln(err)
	}

	t := make(theme)

	t.openDefaultTheme(dist)

	var posts []post.Post
	var items []rss.Item

	for _, md := range mds {
		p := post.New()
		if err := p.Parse(md); err == nil {
			posts = append(posts, p)
			link := fmt.Sprintf(`%s/post/%s.html`, c.URL, p.Slug)
			i := rss.Item{
				Title:       p.Title,
				Link:        link,
				Description: sliceContent(p.Content),
				PubDate:     p.Date.Format(time.ANSIC),
				GUID:        p.Slug,
			}
			items = append(items, i)
		}
	}

	for _, p := range posts {
		buildPost(t, dist, p, c)
	}

	buildIndex(t, dist, indexParams{c.Title, c.Description, posts})
	buildStyle(t, dist)

	ch := rss.Channel{Title: c.Title, Description: c.Description, Link: c.URL, Generator: "", Items: items}
	buildFeed(dist, ch)
}

func buildIndex(t theme, dist string, data indexParams) error {
	tmpl := t.parse("classic", "index")
	f, err := os.Create(path.Join(dist, "index.html"))
	if err != nil {
		return err
	}

	defer f.Close()

	tmpl.Execute(f, data)
	return nil
}

func buildPost(t theme, dist string, p post.Post, c config.Conf) {
	filename := fmt.Sprintf("%s.html", p.Slug)
	path := path.Join(dist, "post", filename)
	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	tmpl := t.parse("", "post")
	tmpl.Execute(f, postParams{p.Title, p.Date, string(p.Content), c.Description})
}

func buildStyle(t theme, dist string) {
	os.WriteFile(path.Join(dist, "style.css"), t["style.css"], perm)
}

func buildFeed(dist string, ch rss.Channel) error {
	f, err := os.Create(path.Join(dist, "feed.rss"))
	if err != nil {
		return err
	}

	defer f.Close()

	rss.Generate(f, ch)

	return nil
}

func sliceContent(s []byte) string {
	r := regexp.MustCompile(`\<\/?\w+\>`)
	ns := r.ReplaceAll(s, []byte{})

	if len(s) > 200 {
		return string(ns)[:200] + "..."
	}

	return string(ns)
}
