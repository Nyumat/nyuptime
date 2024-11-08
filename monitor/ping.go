// Service monitor checks if a website is up or down
package monitor

import (
	"context"
	"net/http"
	"strings"
)

// PingResponse is the response from the Ping endpoint
type PingResponse struct {
	Up bool `json:"up"`
}

// Pings a specific site and determines whether it's up or down during at request time
//encore:api public path=/ping/*url
func Ping(ctx context.Context, url string) (*PingResponse, error) {
    // If the url does not start with "http:" or "https:", default to "https:".
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false}, nil
	}
	resp.Body.Close()

    // 2xx and 3xx status codes are considered up
    up := resp.StatusCode < 400
    return &PingResponse{Up: up}, nil
}
