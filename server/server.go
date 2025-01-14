package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

func handleError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	if status == 404 {
		fmt.Fprintf(w, message)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}
func handleRegister(w http.ResponseWriter, r *http.Request) {

}
