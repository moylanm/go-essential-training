// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct == "" {
		return "", fmt.Errorf("no content type")
	}

	return ct, nil
}
