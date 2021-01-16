package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type result struct {
	Num        int
	Year       string
	Img        string
	Title, Alt string
	Transcript string
}

const (
	cacheFilename = "xkcd.cache"
	APIURL        = "https://xkcd.com/%d/info.0.json"
)

var (
	buildCache = flag.Bool("build-cache", false, "Download and store the cache from the xkcd API")
	comic      = flag.Int("comic", 0, "Show info about this comic")
)

func main() {
	flag.Parse()

	switch {
	case *buildCache:
		fmt.Println("downloading infos")
		err := downloadInfos()
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't download info: %v", err)
			os.Exit(1)
		}
	case 0 < *comic:
		fmt.Println("showing comic")
		err := showComic(*comic)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't show comic %v: %v", *comic, err)
			os.Exit(1)
		}
	case 1 < len(os.Args):
		fmt.Printf("searching for terms: %v\n", os.Args[1:])
		searchComics(os.Args[1:])
	default:
		fmt.Println("Usage: ./xkcd [terms] or ./xkcd -comic n or ./xkcd -build-cache")
	}
}

func downloadInfos() error {
	f, err := os.Create(cacheFilename)
	if err != nil {
		return err
	}
	defer f.Close()

L:
	for i := 1; true; i++ {
		if 404 == i {
			continue
		}

		fmt.Println(i)
		ok, err := downloadInfo(f, i)
		switch {
		case err != nil:
			return err
		case !ok:
			break L
		}
	}
	return nil
}

func downloadInfo(w io.Writer, id int) (bool, error) {
	resp, err := http.Get(fmt.Sprintf(APIURL, id))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return false, nil
	default:
		return false, fmt.Errorf("query failed: %v", resp.StatusCode)
	}

	io.Copy(w, resp.Body)
	return true, nil
}

func showComic(n int) error {
	f, err := os.Open(cacheFilename)
	if err != nil {
		return err
	}

	var res result
	dec := json.NewDecoder(f)
	for {
		err = dec.Decode(&res)
		switch {
		case err != nil:
			return err
		case n == res.Num:
			fmt.Println(res)
			return nil
		}
	}
}

func searchComics(terms []string) error {
	f, err := os.Open(cacheFilename)
	if err != nil {
		return err
	}

	var res result
	dec := json.NewDecoder(f)
L:
	for {
		if err = dec.Decode(&res); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		for _, t := range terms {
			if -1 == strings.Index(res.Transcript, t) {
				continue L
			}
		}
		fmt.Printf("%v\n\n", res)
	}
}
