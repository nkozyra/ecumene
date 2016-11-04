package ecumene

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {

	r := mux.NewRouter()
	r.HandleFunc("/address", AddressHandler)

	http.ListenAndServe(Port, r)
}
