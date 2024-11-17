package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const (
	contentTypeNDJSON = "application/x-ndjson"
	contentTypeJSON   = "application/json"
)

var addr = flag.String("addr", ":9999", "addr to listen")

func main() {

	// serve json object(s) only containing uuidv4 for now, with appropriate header.
	// optional query param `count` generates multiple or potentially 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		c, err := strconv.ParseUint(r.URL.Query().Get("count"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if c == 1 {
			w.Header().Set("Content-Type", contentTypeJSON)
			fmt.Fprintf(w, "{\"uuidv4\":\"%s\"}", uuid.NewString())
			return
		}

		w.Header().Set("Content-Type", contentTypeNDJSON)
		for ; c > 0; c-- {
			fmt.Fprintf(w, "{\"uuidv4\":\"%s\"}\n", uuid.NewString())
		}
	})

	log.Println("attempting to listen on " + *addr)

	// you are expected to be able to read stack traces
	log.Fatal(http.ListenAndServe(*addr, nil))
}
