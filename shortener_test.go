package shortener

import (
	"testing"
)

var uri []string

func init() {
	uri = []string{
		`https://similar-sites.com`,
		`https://similar-sites.com`,
		`https://other-sites.com`,
	}
}

func TestCreateShortLink(t *testing.T) {
	shortLinks := createShortLinks()
	for _, shortLink := range shortLinks {
		if shortLink == "" {
			t.Error("Could not create short link", shortLink)
		} else {
			t.Log("Success create short link", shortLink)
		}
	}

	if shortLinks[0] != shortLinks[1] {
		t.Error("Short link not generated correctly", shortLinks[0], shortLinks[1])
	}
	if shortLinks[2] == shortLinks[0] || shortLinks[2] == shortLinks[1] {
		t.Error("All links are identical", shortLinks)
	}
}

func TestGetUrlByShortLink(t *testing.T) {
	shortLinks := createShortLinks()
	for index, shortLink := range shortLinks {
		origUrl, ok := GetUrlByShortLink(shortLink)
		if !ok {
			t.Error("No original URI found by short link", origUrl, uri[index])
		}
		if origUrl != uri[index] {
			t.Error("Error getting original URI via short link", origUrl, uri[index])
		}
	}
}

func BenchmarkCreateShortLink(b *testing.B) {
	b.StartTimer()
	createShortLinks()
	b.StopTimer()
}

func BenchmarkGetUrlByShortLink(b *testing.B) {
	shortLinks := createShortLinks()
	b.StartTimer()
	for _, shortLink := range shortLinks {
		GetUrlByShortLink(shortLink)
	}
	b.StopTimer()
}

func createShortLinks() []string {
	var shortLinks []string
	for _, link := range uri {
		shortLinks = append(shortLinks, CreateShortLink(link))
	}

	return shortLinks
}
