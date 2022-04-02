package net

import (
	"evelp/log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var client = &http.Client{Timeout: 10 * time.Second}

var backoffSchedule = []time.Duration{
	1 * time.Second,
	2 * time.Second,
	3 * time.Second,
}

func GetWithRetries(request string) (*http.Response, error) {
	var resp *http.Response
	var err error

	for _, backoff := range backoffSchedule {
		resp, err = Get(request)
		code := resp.StatusCode

		if code == http.StatusNotFound {
			return nil, err
		}

		if err == nil {
			if code == http.StatusOK {
				break
			}
			err = errors.Errorf("http request %s error status code %d", request, code)
		}

		log.Debugf("http request %s failed: %+v \nretrying in %v", request, err.Error(), backoff)
		time.Sleep(backoff)
	}

	if err != nil {
		return nil, errors.WithMessage(err, "http all request retries failed")
	}

	return resp, nil
}

func Get(request string) (*http.Response, error) {
	resp, err := client.Get(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}
