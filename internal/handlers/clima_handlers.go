package handlers

import "net/http"

func ClimaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Clima Handler"))
}
