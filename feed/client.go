package feed

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/cretz/bine/tor"
	"github.com/ipsn/go-libtor"
)

// TorClient ...
type TorClient struct {
	*http.Client
	tor        *tor.Tor
	dialCancel context.CancelFunc
}

// NewTorClient ...
func NewTorClient() *TorClient {
	t, err := tor.Start(nil, &tor.StartConf{ProcessCreator: libtor.Creator, DebugWriter: os.Stderr})
	if err != nil {
		panic(err)
	}

	// Wait at most a minute to start network and get
	dialCtx, dialCancel := context.WithTimeout(context.Background(), time.Minute)

	// Make connection
	dialer, err := t.Dialer(dialCtx, nil)
	if err != nil {
		panic(err)
	}

	return &TorClient{
		dialCancel: dialCancel,
		tor:        t,
		Client: &http.Client{
			Transport: &http.Transport{DialContext: dialer.DialContext},
			Timeout:   30 * time.Second,
		},
	}
}

// Close ...
func (t *TorClient) Close() {
	t.tor.Close()
	t.dialCancel()
}

// NewGenericClient ...
func NewGenericClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
