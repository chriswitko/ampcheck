package link

import (
    "testing"
)

var testURLs = []string{
	"https://ymotongpoo.github.io/demos/amp/amp.html",
}

func TestValidate(t *testing.T) {
    for _, u := range testURLs {
        links, err := Validate(u)
        if err != nil {
            t.Fatalf("%v", err)
        }
        if !links.Valid {
            t.Fatalf("canonical: %v, amphtml: %v", links.Canonical, links.AMP)
        }
    }
}