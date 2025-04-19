package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch[T interface{}](url string) (T, error) {
	var zero T

	if data, ok := GlobalCache.Get(url); ok {
		return unmarshal[T](data)
	}

	res, err := http.Get(url)
	if err != nil {
		return zero, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("expected status OK, got %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, err
	}

	GlobalCache.Add(url, data)

	return unmarshal[T](data)
}

func unmarshal[T interface{}](data []byte) (T, error) {
	var out T
	if err := json.Unmarshal(data, &out); err != nil {
		return out, err
	}
	return out, nil
}
