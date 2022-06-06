package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

const perm = 0444

type indexParams struct {
	Title       string
	Description string
	Posts       []post
}

type postParams struct {
	Title       string
	Date        time.Time
	Content     string
	Description string
}

func build(c conf, dist string) {
	os.RemoveAll(dist)
	if err := os.MkdirAll(path.Join(dist, "post"), 0777); err != nil {
		log.Fatalln(err)
	}

	mds, err := getPosts()
	if err != nil {
		log.Fatalln(err)
	}

	t := make(theme)

	t.openDefaultTheme()

	var posts []post
	var items []item

	for _, md := range mds {
		p := newPost()
		if err := p.parse(md); err == nil {
			posts = append(posts, p)
			link := fmt.Sprintf(`%s/post/%s.html`, c.URL, p.Slug)
			i := item{p.Title, link, sliceContent(p.Content), p.Date.Format(time.ANSIC), p.Slug}
			items = append(items, i)
		}
	}

	for _, p := range posts {
		buildPost(t, dist, p, c)
	}

	buildIndex(t, dist, indexParams{c.Title, c.Description, posts})
	buildStyle(t, dist)

	ch := channel{c.Title, c.Description, c.URL, "", items}
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

func buildPost(t theme, dist string, p post, c conf) {
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

func buildFeed(dist string, ch channel) error {
	f, err := os.Create(path.Join(dist, "feed.rss"))
	if err != nil {
		return err
	}

	defer f.Close()

	generateFeed(f, ch)

	return nil
}

func sliceContent(s []byte) string {
	if len(s) > 200 {
		return string(s)[:200] + "..."

	}
	return string(s)
}
