package urlShort

import (
	"crypto/md5"
	"encoding/hex"
	storage "github.com/neimus/go-short-url/storage/memory"
	"math/rand"
	"net/url"
	"time"
)

const URL_SCHEME string = "https://"
const MAIN_URL string = "short.lnk"
const URL_SEPARATOR string = "/"
const LETTER = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LEN_SHORT_LINK int = 7

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateShortLink(uri string) string {
	hash := createHash(uri)
	urlType, ok := storage.GetUrlByHash(hash)
	if !ok {
		shortLink := generateShortLink()
		urlType = storage.Save(hash, uri, shortLink)
	}

	return URL_SCHEME + MAIN_URL + urlType.ShortLink
}

func GetUrlByShortLink(uriShortLink string) (string, bool) {
	u, errParse := url.Parse(uriShortLink)
	if errParse == nil {
		urlType, ok := storage.GetUrlByShortLinks(u.RequestURI())
		if ok {
			return urlType.Uri, true
		}
	}
	return "", false
}

func generateShortLink() string {
	b := make([]byte, LEN_SHORT_LINK)
	for i := range b {
		b[i] = LETTER[rand.Intn(len(LETTER))]
	}
	shortLink := URL_SEPARATOR + string(b)

	if storage.HasShortLink(shortLink) {
		return generateShortLink()
	}

	return shortLink
}

func createHash(uri string) string {
	hash := md5.Sum([]byte(uri))
	return hex.EncodeToString(hash[:])
}
