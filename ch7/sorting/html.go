package main

import (
	"html/template"
	"log"
	"net/http"
)

var trackTemplate = template.Must(template.New("tracklist").Parse(`
<table>
 <tr>
  <td><a href="/name">name</a></td>
  <td><a href="/artist">artist</a></td>
  <td><a href="/album">album</a></td>
  <td><a href="/year">year</a></td>
  <td><a href="/length">length</a></td>
 </tr>
{{range .Tracks}}
 <tr>
  <td>{{.Name}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
 </tr>
{{end}}
</table>
`))

type tableHandler struct {
	ms *multiSorter
}

func (h *tableHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/name":
		h.ms.sortBy(Name)
	case "/artist":
		h.ms.sortBy(Artist)
	case "/album":
		h.ms.sortBy(Album)
	case "/year":
		h.ms.sortBy(Year)
	case "/length":
		h.ms.sortBy(Length)
	}

	if err := trackTemplate.Execute(w, h.ms); err != nil {
		log.Fatal(err)
	}
}

var _ http.Handler = &tableHandler{}
