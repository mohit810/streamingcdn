package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mohit810/streamingcdn/signal"
)

//func init() {
//	// Generate pem file for https
//	signal.GenPem()
//}

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	r := httprouter.New()
	signal.HTTPSDPServer(r)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	if err != nil {
		panic(err)
	}
	//panic(http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert.pem", "key.pem", r))
}
