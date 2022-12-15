package app

import (
	"fmt"
	"math/rand"
)

type Services interface {
	ShortenUrl(url string) (shortenedUrl string)
	RetrieveOriginalUrl(shortUrl string) (url string, err error)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randStringBytes - Create random short link
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var urlMap = make(map[string]string)

var revUrlMap = make(map[string]string)

func (s *service) ShortenUrl(url string) (shortenedUrl string) {

	shortenedUrl = urlMap[url]

	if shortenedUrl == "" {
	SHORT_URL:
		shortenedUrl = randStringBytes(10)

		if revUrlMap[shortenedUrl] != "" {
			goto SHORT_URL
		}

		revUrlMap[shortenedUrl] = url
		urlMap[url] = shortenedUrl

	}
	return
}

func (s *service) RetrieveOriginalUrl(shortUrl string) (url string, err error) {

	url = revUrlMap[shortUrl]

	if url == "" {
		err = fmt.Errorf("url not found")
	}
	return
}
