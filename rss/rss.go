package rss

import (
	"encoding/xml"
	"io"
)

type RSS struct {
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Generator   string `xml:"generator"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func Generate(w io.Writer, c Channel) error {
	encoder := xml.NewEncoder(w)
	c.Generator = "rramiachraf/static"
	data := RSS{"2.0", c}

	err := encoder.Encode(&data)
	if err != nil {
		return err
	}

	return nil
}
