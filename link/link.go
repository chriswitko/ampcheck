package link

import (
    "bytes"
	"net/http"
	"reflect"

    "golang.org/x/net/html"
	"launchpad.net/xmlpath"
)

var (
	canonical = xmlpath.MustCompile("/html/head/link[@rel='canonical']/@href")
	amphtml   = xmlpath.MustCompile("/html/head/link[@rel='amphtml']/@href")
)

type Links struct {
	Canonical string
	AMP       string
	Valid     bool
}

// Validate link relations.
func Validate(urlStr string) (*Links, error) {
	src, err := parse(urlStr)
	if err != nil {
		return src, err
	}
	l1, err := parse(src.Canonical)
	if err != nil {
		return l1, err
	}
	l2, err := parse(src.AMP)
	if err != nil {
		return l2, err
	}
	valid := reflect.DeepEqual(l1, l2)
	src.Valid = valid
	return src, nil
}

func parse(urlStr string) (*Links, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
    srcRoot, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }
    var b bytes.Buffer
    html.Render(&b, srcRoot)
	root, err := xmlpath.ParseHTML(bytes.NewReader(b.Bytes()))
	if err != nil {
		return nil, err
	}
	links := Links{}
	if canonicalURL, ok := canonical.String(root); ok {
		links.Canonical = canonicalURL
	}
	if ampURL, ok := amphtml.String(root); ok {
		links.AMP = ampURL
	}
	return &links, nil
}
