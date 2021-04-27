package sexpr

import (
	"fmt"
	"os"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var strangelove Movie

func init() {
	strangelove = Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peters Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
}

func TestMarshal(t *testing.T) {
	bs, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf(err.Error())
	}
	os.Stdout.Write(bs)
}

func TestUnmarshal(t *testing.T) {
	bs, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var copy Movie
	Unmarshal(bs, &copy)
	fmt.Println(copy)
}
