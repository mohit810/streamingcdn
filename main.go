package main

import (
	"flag"
	"github.com/rs/cors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mohit810/streamingcdn/signal"
)

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	r := httprouter.New()
	signal.HTTPSDPServer(r)
	handler := cors.Default().Handler(r)
	panic(http.ListenAndServe(":"+strconv.Itoa(*port), handler))
}
