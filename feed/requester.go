package feed

import (
	"fmt"
	"net/http"
)

// NewRequester ..
func NewRequester(httpClient *http.Client, agent string) Requester {
	return &MyRequester{
		Client: httpClient,
		agent:  agent,
	}
}

// MyRequester ...
type MyRequester struct {
	*http.Client
	agent string
}

// Request ...
func (f *MyRequester) Request(url string) (*http.Response, error) {

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", f.agent)

	// Make request
	res, err := f.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 && res.StatusCode != 304 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res, nil

}
