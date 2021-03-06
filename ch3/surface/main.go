// Surface computes an SVG redering of a 3D surface function.
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/surface", handleRenderSVG)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
