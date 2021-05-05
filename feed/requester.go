package feed

import (
	"fmt"
	"net/http"

	"github.com/corpix/uarand"
)

// NewRequester ..
func NewRequester(httpClient *http.Client) Requester {
	return &MyRequester{
		Client:  httpClient,
		counter: 0,
	}
}

// MyRequester ...
type MyRequester struct {
	*http.Client
	counter int
	agent   string
}

// Request ...
func (f *MyRequester) Request(url string) (*http.Response, error) {

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	f.updateAgent()
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

func (f *MyRequester) updateAgent() {
	if f.agent == "" || f.counter >= 10 {
		f.agent = uarand.GetRandom()
		f.counter = 0
	}
	f.counter++
}
