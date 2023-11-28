package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpClient struct {
	Client     *http.Client
	MaxRetries int
	RetryDelay time.Duration
	TimeOut    time.Duration
}

func NewHTTPClient(maxRetries int, retryInterval time.Duration, timeout time.Duration) *HttpClient {
	return &HttpClient{
		Client:     &http.Client{},
		MaxRetries: maxRetries,
		RetryDelay: retryInterval,
		TimeOut:    timeout,
	}
}

func (hc *HttpClient) doRequest(req *http.Request) (resp *http.Response, err error) {
	i := 0
	for ; i < hc.MaxRetries; i++ {
		resp, err = hc.Client.Do(req)
		if err != nil {
			time.Sleep(hc.RetryDelay)
			continue
		}
		return
	}
	if i == hc.MaxRetries {
		err = fmt.Errorf("max retries exceeded")
	}

	return
}

func (hc *HttpClient) Get(c context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(c, hc.TimeOut)
	defer cancel()
	req = req.WithContext(ctx)
	return hc.doRequest(req)
}
