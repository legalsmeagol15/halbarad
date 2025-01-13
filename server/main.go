package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/*.js", handleJS)

	err := http.ListenAndServe(":9001", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server is closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
