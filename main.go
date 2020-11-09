package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"practice/streamingcdn/signal"
	"strconv"
)

func main() {
	port := flag.Int("port", 8000, "http server port")
	flag.Parse()
	r := httprouter.New()
	signal.HTTPSDPServer(r)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	if err != nil {
		panic(err)
	}
}
