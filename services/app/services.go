package app

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

type URL struct {
	u      string
	expiry time.Time
}

var expiryMap = make(map[string]URL)

var urlMap = make(map[string]string)

var revUrlMap = make(map[string]string)

var expiryInterval = time.Minute * 1

func (s *service) ShortenUrl(url string) (shortenedUrl string) {

	if url == "" {
		s.Log.Errorln("url is empty")
		return
	}

	shortenedUrl = urlMap[url]

	if shortenedUrl == "" {
	SHORT_URL:
		shortenedUrl = randStringBytes(10)

		if revUrlMap[shortenedUrl] != "" {
			goto SHORT_URL
		}

		revUrlMap[shortenedUrl] = url
		urlMap[url] = shortenedUrl

		expiryMap[shortenedUrl] = URL{
			u:      shortenedUrl,
			expiry: time.Now().Add(expiryInterval),
		}

	}
	return
}

func (s *service) RetrieveOriginalUrl(shortUrl string) (url string, err error) {

	expiry_ := expiryMap[shortUrl]

	if time.Now().After(expiry_.expiry) {
		err = fmt.Errorf("shortned url has expired at %v", expiry_.expiry)
		return
	}

	url = revUrlMap[shortUrl]

	if url == "" {
		err = fmt.Errorf("url %s  not found", shortUrl)
		return
	}

	if !strings.Contains(url, "http") {
		url = fmt.Sprintf("https://%s", url)
	}
	return
}
