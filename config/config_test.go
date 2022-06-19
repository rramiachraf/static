package config

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestParse(t *testing.T) {
	title := "example"

	path := path.Join(t.TempDir(), "config.yml")
	content := fmt.Sprintf("title: %s\n", title)

	err := os.WriteFile(path, []byte(content), 0777)
	if err != nil {
		t.Fatalf("[FAIL] cannot create configuration file")
	}

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("[FAIL] error parsing file (%s)", err)
	}

	if c.Title != title {
		t.Fatalf("[FAIL] expected %s, got %s", title, c.Title)
	}
}
