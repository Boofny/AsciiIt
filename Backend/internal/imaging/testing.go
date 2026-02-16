package main

import (
	"fmt"
	"net/http"
	"strings"
)

const URLs = "https://picsum.photos/200/300"

func main() {
	resp, err := http.Head(URLs)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rr := resp.Header

	fmt.Println(rr)
	r := "https://ichef.bbci.co.uk/ace/standard/976/cpsprodpb/14235/production/_100058428_mediaitem100058424.jpg"

	t, err := isImageURL(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}

func isImageURL(url string) (bool, error) {
	resp, err := http.Head(url) // Use HEAD request to get headers only
	if err != nil {
		return false, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("non-OK HTTP status: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	fmt.Println(contentType)
	// Check if the Content-Type starts with "image/"
	if strings.HasPrefix(contentType, "image/") {
		return true, nil
	}

	return false, nil
}
