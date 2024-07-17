package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

// options contains the configuration for the webserver
type options struct {
	path   string
	listen string
}

func main() {
	var opts options
	flag.StringVar(&opts.path, "path", "/", "path to serve")
	flag.StringVar(&opts.listen, "listen", ":8080", "the address to listen on")
	flag.Parse()

	http.HandleFunc(opts.path, func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "ok"); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("listening on %s for requests to %s", opts.listen, opts.path)
	log.Fatal(http.ListenAndServe(opts.listen, nil))
}
