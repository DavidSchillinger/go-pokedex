package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CachedFetch[T any](url string, out *T) error {
	if data, ok := GlobalCache.Get(url); ok {
		return unmarshal(data, out)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status OK, got %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	GlobalCache.Add(url, data)

	return unmarshal(data, out)
}

func unmarshal[T any](data []byte, out *T) error {
	if err := json.Unmarshal(data, &out); err != nil {
		return err
	}
	return nil
}
