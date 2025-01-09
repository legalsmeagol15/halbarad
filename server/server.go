package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Starting backend server...")
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
}

func NewServer(
	logger *Logger,
	config *Config,
	commentStore *commentStore,
	anotherStore *anotherStore,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		Logger,
		Config,
		commentStore,
		anotherStore,
	)
	var handler http.Handler = mux
	handler = someMiddleware(handler)
	handler = someMiddleware2(handler)
	handler = someMiddleware3(handler)
	return handler
}
