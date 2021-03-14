package main

import (
	"sort"
)

type sortKey int

const (
	Name sortKey = iota
	Artist
	Album
	Year
	Length
)

type multiSorter struct {
	Tracks   []*Track
	lastKeys []sortKey
}

func makeMultiSorter(ts []*Track) *multiSorter {
	nts := make([]*Track, len(ts))
	copy(nts, ts)

	return &multiSorter{Tracks: nts}
}

func (m *multiSorter) Len() int {
	return len(m.Tracks)
}

func (m *multiSorter) Less(i, j int) bool {
	for _, key := range m.lastKeys {
		switch {
		case Name == key && m.Tracks[i].Name != m.Tracks[j].Name:
			return m.Tracks[i].Name < m.Tracks[j].Name
		case Artist == key && m.Tracks[i].Artist != m.Tracks[j].Artist:
			return m.Tracks[i].Artist < m.Tracks[j].Artist
		case Album == key && m.Tracks[i].Album != m.Tracks[j].Album:
			return m.Tracks[i].Album < m.Tracks[j].Album
		case Year == key && m.Tracks[i].Year != m.Tracks[j].Year:
			return m.Tracks[i].Year < m.Tracks[j].Year
		case Length == key && m.Tracks[i].Length != m.Tracks[j].Length:
			return m.Tracks[i].Length < m.Tracks[j].Length
		}
	}
	return false
}

func (m *multiSorter) Swap(i, j int) {
	m.Tracks[i], m.Tracks[j] = m.Tracks[j], m.Tracks[i]
}

func (m *multiSorter) sortBy(s sortKey) {
	m.lastKeys = append([]sortKey{s}, m.lastKeys...)
	sort.Sort(m)
}

var _ sort.Interface = (*multiSorter)(nil)
