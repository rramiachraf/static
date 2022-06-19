package highlight

import (
	"bytes"
	"regexp"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func format(lang, code string) ([]byte, error) {
	var lexer chroma.Lexer

	if lang == "" {
		lexer = lexers.Analyse(code)
	}

	if lang != "" {
		lexer = lexers.Get(lang)
	}

	if lexer == nil {
		lexer = lexers.Fallback
	}

	lexer = chroma.Coalesce(lexer)

	style := styles.Get("xcode-dark")
	if style == nil {
		style = styles.Fallback
	}

	formatter := html.New(html.WithLineNumbers(true))

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return nil, err
	}

	var w bytes.Buffer

	err = formatter.Format(&w, style, iterator)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func ParseSnippet(snippet []byte) (string, string) {
	r := regexp.MustCompile(`(?ms)\x60\x60\x60(\w+)?\s?(.*?)\x60\x60\x60`)
	m := r.FindSubmatch(snippet)

	var lang string
	var code string

	if len(m) == 3 {
		lang = string(m[1])
		code = string(m[2])
	}

	if len(m) == 2 {
		code = string(m[2])
	}

	return lang, code
}

func Highlight(content []byte) []byte {
	r := regexp.MustCompile(`(?ms)\x60\x60\x60.*?\x60\x60\x60`)

	return r.ReplaceAllFunc(content, func(s []byte) []byte {
		lang, code := ParseSnippet(s)
		if highlighted, err := format(lang, code); err == nil {
			return highlighted
		}
		return s
	})
}
