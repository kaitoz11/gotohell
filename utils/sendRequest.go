package utils

import (
	"bytes"
	"log"
	"net/http"
)

func SendGET(target string, cookies []*http.Cookie, headers map[string][]string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header = headers

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	return client.Do(req)
}

func SendPOST(target string, body string, cookies []*http.Cookie, headers map[string][]string) (*http.Response, error) {
	// jsonBody := []byte(body)
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPost, target, bodyReader)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header = headers

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	return client.Do(req)
}
