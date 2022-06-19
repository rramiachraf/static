package rss

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/gosimple/slug"
)

func TestGenerate(t *testing.T) {
	filename := path.Join(t.TempDir(), "feed.xml")

	f, err := os.Create(filename)
	if err != nil {
		t.Fatalf("[FAIL] %s", err)
	}

	title := "John Doe's feed"
	description := "My RSS feed is better than yours."
	link := "https://example.com"
	generator := "rramiachraf/static"

	itemTitle := "Doctors need to work on their handwriting"
	itemLink := fmt.Sprintf("https://example.com/post/%s", slug.Make(itemTitle))
	itemDescription := "Doctors, your handwriting is horrible"
	itemPubDate := "Sun Jun 19 18:46:00 2022"
	itemGUID := slug.Make(itemLink)

	i := Item{itemTitle, itemLink, itemDescription, itemPubDate, itemGUID}
	c := Channel{title, description, link, generator, []Item{i}}

	if err = Generate(f, c); err != nil {
		t.Fatalf("[FAIL] %s", err)
	}

	f.Close()

	expected, err := xml.Marshal(RSS{"2.0", c})
	if err != nil {
		t.Fatalf("[FAIL] %s", err)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("[FAIL] %s", err)
	}

	if identical := bytes.Compare(b, expected); identical != 0 {
		t.Fatalf("[FAIL] expected %s got %s", expected, b)
	}
}
