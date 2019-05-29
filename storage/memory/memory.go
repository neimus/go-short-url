package memory

import (
	types "github.com/neimus/go-short-url/types"
)

var urls map[string]types.Url
var shortLinks map[string]string

func init() {
	urls = make(map[string]types.Url)
	shortLinks = make(map[string]string)
}

func Save(hash string, uri string, shortLink string) types.Url {
	urls[hash] = types.Url{Uri: uri, ShortLink: shortLink}
	shortLinks[shortLink] = hash

	return urls[hash]
}

func GetUrlByHash(hash string) (types.Url, bool) {
	urlType, ok := urls[hash]
	return urlType, ok
}

func GetUrlByShortLinks(shortLink string) (types.Url, bool) {
	hash, ok := shortLinks[shortLink]
	if ok {
		urlType, isset := urls[hash]
		return urlType, isset
	}

	return types.Url{}, false
}

func HasShortLink(shortLink string) bool {
	_, ok := shortLinks[shortLink]
	return ok
}
