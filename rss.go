package main

import (
	"encoding/xml"
	"io"
)

type rss struct {
	Version string  `xml:"version,attr"`
	Channel channel `xml:"channel"`
}

type channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Generator   string `xml:"generator"`
	Item        []item `xml:"item"`
}

type item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func generateFeed(w io.Writer, c channel) error {
	encoder := xml.NewEncoder(w)
	c.Generator = "rramiachraf/static"
	data := rss{"2.0", c}

	err := encoder.Encode(&data)
	if err != nil {
		return err
	}

	return nil
}
