package main

import (
	"flag"
	"github.com/rs/cors"
	"log"
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
	log.Println("Server is Up and Running at Port:" + strconv.Itoa(*port))
	handler := cors.Default().Handler(r)
	panic(http.ListenAndServe(":"+strconv.Itoa(*port), handler))
}
