package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	logger log.Logger
)

func main() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetOutput(os.Stdout)
	logger.Println("Starting web server...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Println("server closed in SIGINT")
			os.Exit(0)
		}
	}()

	// Don't use the DefaultServeMux because it is a global variable that is accessible to any code.
	mux := http.NewServeMux()

	// TODO:  prepend the canonical host name
	root_fs := http.FileServer(http.Dir("../web/"))
	mux.Handle("GET /", root_fs)
	mux.HandleFunc("POST /login", handleLogin)

	// Do the business logic of handling requests
	err := http.ListenAndServe(":9001", mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Println("server closed normally")
	} else if err != nil {
		logger.Fatalf("error starting server: %s\n", err)
	}

}
