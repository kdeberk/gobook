package main

import (
	"reflect"
	"testing"
)

func TestMultiSorter(t *testing.T) {
	ts := []*Track{
		{Name: "Pow R. Toc. H", Album: "The Piper At The Gates Of Dawn", Artist: "Pink Floyd", Year: 1967, Length: length("4m26s")},
		{Name: "Fear of Flying", Album: "Beat", Artist: "Bowery Electric", Year: 1996, Length: length("5m39s")},
		{Name: "I Dreamed I Saw St. Augustine", Album: "John Wesley Harding", Artist: "Bob Dylan", Year: 1967, Length: length("3m53s")},
		{Name: "Bold As Love", Album: "Axis: Bold As Love", Artist: "Jimi Hendrix", Year: 1967, Length: length("4m11s")},
		{Name: "Interstellar Overdrive", Album: "The Piper At The Gates Of Dawn", Artist: "Pink Floyd", Year: 1967, Length: length("9m41s")},
	}

	tts := []struct {
		name string
		keys []sortKey
		exp  []*Track
	}{
		{name: "only by name",
			keys: []sortKey{Name},
			exp:  []*Track{ts[3], ts[1], ts[2], ts[4], ts[0]},
		},
		{name: "year then name",
			keys: []sortKey{Name, Year},
			exp:  []*Track{ts[3], ts[2], ts[4], ts[0], ts[1]},
		},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ms := makeMultiSorter(ts)
			for _, key := range tt.keys {
				ms.sortBy(key)
			}

			if !reflect.DeepEqual(tt.exp, ms.Tracks) {
				t.Errorf("Expected order different from actual.\n\tbexp: %v\n\tact: %v", tt.exp, ms.Tracks)
			}
		})
	}
}
