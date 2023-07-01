package answer

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func Test_GoQuery(t *testing.T) {
	html := `<html><head><title>First parse</title></head></html>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(doc.Find("title").Text())
}
