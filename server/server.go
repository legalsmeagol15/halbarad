package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleIndex")

}
func handleJS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleJS")

}
