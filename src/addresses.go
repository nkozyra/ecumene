package ecumene

import (
	"encoding/json"
	"fmt"
	"net/http"

	parser "github.com/openvenues/gopostal/parser"
)

func ParseAddress(ad string) []parser.ParsedComponent {
	parsed := parser.ParseAddress(ad)

	return parsed
}

func AddressHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("address")
	b := ParseAddress(a)
	c, _ := json.Marshal(&b)
	fmt.Fprintln(w, string(c))

}
