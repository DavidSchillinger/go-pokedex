package main

import (
	"fmt"
	"io"
	"net/http"
)

func CachedFetch(url string) ([]byte, error) {
	if data, ok := GlobalCache.Get(url); ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status OK, got %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	GlobalCache.Add(url, data)

	return data, nil
}
