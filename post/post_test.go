package post

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	dir := t.TempDir()

	var files []string

	for _, f := range files {
		p := path.Join(dir, f)
		files = append(files, p)

		if err := os.WriteFile(p, []byte{}, 0777); err != nil {
			continue
		}
	}

	results, err := GetAll(dir)
	if err != nil {
		t.Fatalf("[FAIL] cannot get md files")
	}

	if !reflect.DeepEqual(results, files) {
		t.Fatalf("[FAIL] expected %v got %v", files, results)
	}
}

func TestParseHeader(t *testing.T) {
	title := "Bald people often have no hair"
	date := "13/06/2022 16:49"

	header := fmt.Sprintf("TITLE %s\nDATE %s\n", title, date)
	results := ParseHeader([]byte(header))

	if results["TITLE"] != title {
		t.Fatalf("[FAIL] expected %s got %s", title, results["TITLE"])
	}

	if results["DATE"] != date {
		t.Fatalf("[FAIL] expected %s got %s", date, results["DATE"])
	}
}

func TestScan(t *testing.T) {
	filename := path.Join(t.TempDir(), "2022.md")

	h := "TITLE title\nDATE date\n"
	c := "Lorem ipsum dolor\nsit amet\n"
	data := fmt.Sprintf("%s\n%s", h, c)

	if err := os.WriteFile(filename, []byte(data), 0777); err != nil {
		t.Fatalf("[FAIL] cannot write to file")
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("[FAIL] cannot open file")
	}

	defer f.Close()

	_, content, err := Scan(f)
	if err != nil {
		t.Fatalf("[FAIL] %s", err)
	}

	if string(content) != c {
		t.Fatalf("[FAIL] expected %q got %q", c, content)
	}
}

func TestParse(t *testing.T) {
	filename := path.Join(t.TempDir(), "go.md")

	title := "An hour, on average, is around 60 minutes"
	date := "18/06/2022 21:16"
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	data := fmt.Sprintf("TITLE %s\nDATE %s\n\n%s", title, date, content)

	if err := os.WriteFile(filename, []byte(data), 0777); err != nil {
		t.Fatalf("[FAIL] cannot create file")
	}

	var p Post
	if err := p.Parse(filename); err != nil {
		t.Fatalf("[FAIL] cannot parse file")
	}

	if p.Title != title {
		t.Fatalf("[FAIL] expected %s got %s", title, p.Title)
	}
}
