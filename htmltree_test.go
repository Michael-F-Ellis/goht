package goht

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	type items struct {
		e   *HtmlTree //what to render
		exp string    // expected result
	}
	table := []items{
		{Html(""), "<html></html>"},
		{P(`class=myclass`), "<p class=myclass></p>"},
		{P(`data-foo="foo text"`), `<p data-foo="foo text"></p>`},
		{Br(``), `<br>`},
		{Null(``), ``},
		{Null(Br(``), Br(``)), `<br><br>`},
	}
	for _, test := range table {
		var b bytes.Buffer
		err := Render(test.e, &b, -1)
		if err != nil {
			t.Errorf("Render failed: %v", err)
		}
		r := b.String()
		if r != test.exp {
			t.Errorf("Expected %s, got %s", test.exp, r)
		}
	}
}

func BenchmarkRender(b *testing.B) {
	meta := Meta(`title="Demo"`)
	head := Head("id=2 class=foo", meta)
	body := Body("id=3 class=bar", Div("", "hello", Br(``)))
	html := Html("", head, body)

	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		_ = Render(html, &b, -1)
	}
}
