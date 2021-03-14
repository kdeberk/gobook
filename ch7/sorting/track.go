package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Track struct {
	Name, Artist, Album string
	Year                int
	Length              time.Duration
}

func (t *Track) String() string {
	return fmt.Sprintf("%v - %v", t.Name, t.Artist)
}

func LoadTracks(r io.Reader, ts []*Track) (err error) {
	return json.NewDecoder(r).Decode(&ts)
}
