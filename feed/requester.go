package feed

import (
	"fmt"
	"net/http"
	"sync"
)

// NewRequester ..
func NewRequester(httpClient *http.Client) Requester {

	return &MyRequester{
		Client: httpClient,
	}
}

// MyRequester ...
type MyRequester struct {
	*http.Client
	lock sync.Mutex
}

// Request ...
func (f *MyRequester) Request(url string) (*http.Response, error) {

	//TODO: throttle

	f.lock.Lock()
	defer f.lock.Unlock()

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0")

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
