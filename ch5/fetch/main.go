package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		f, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url: %v", err)
			continue
		}
		fmt.Printf("%v -> %v: %v\n", url, f, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err != nil {
			err = fmt.Errorf("when closing: %w", closeErr)
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
