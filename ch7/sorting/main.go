package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	ts := []*Track{
		{Name: "Pow R. Toc. H", Album: "The Piper At The Gates Of Dawn", Artist: "Pink Floyd", Year: 1967, Length: length("4m26s")},
		{Name: "Fear of Flying", Album: "Beat", Artist: "Bowery Electric", Year: 1996, Length: length("5m39s")},
		{Name: "I Dreamed I Saw St. Augustine", Album: "John Wesley Harding", Artist: "Bob Dylan", Year: 1967, Length: length("3m53s")},
		{Name: "Bold As Love", Album: "Axis: Bold As Love", Artist: "Jimi Hendrix", Year: 1967, Length: length("4m11s")},
		{Name: "Interstellar Overdrive", Album: "The Piper At The Gates Of Dawn", Artist: "Pink Floyd", Year: 1967, Length: length("9m41s")},
	}

	http.Handle("/", &tableHandler{ms: makeMultiSorter(ts)})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
