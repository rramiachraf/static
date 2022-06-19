package highlight

import (
	"fmt"
	"testing"
)

const lang = "ruby"
const code = "puts \"Hello World!\""

var content = fmt.Sprintf("Lorem ipsum dolor sit amet\n ```%s\n%s```", lang, code)

func TestParseSnippet(t *testing.T) {
	l, c := ParseSnippet([]byte(content))

	if l != lang {
		t.Errorf("[FAIL] expected %s got %s", lang, l)
	}

	if c != code {
		t.Errorf("[FAIL] expected %s got %s", code, c)
	}
}
