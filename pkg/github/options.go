package github

import (
	"fmt"
	"net/http"
)

type GithubClientOption func(*GithubClient)

func WithHttpClient(client *http.Client) GithubClientOption {
	return func(gc *GithubClient) {
		gc.httpClient = client
	}
}

type transport struct {
	token string
}

func (t transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return http.DefaultTransport.RoundTrip(req)
}

func WithToken(token string) GithubClientOption {
	return func(gc *GithubClient) {
		gc.httpClient.Transport = transport{token: token}
	}
}
