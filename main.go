package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"github.com/mohit810/streamingcdn/signal"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

/*func init() { //uncomment for cloud testing
	// Generate pem file for https
	signal.GenPem()
}*/

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	r := httprouter.New()
	signal.HTTPSDPServer(r)
	handler := cors.Default().Handler(r)
	panic(http.ListenAndServe(":"+strconv.Itoa(*port), handler))
	//panic(http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert.pem", "key.pem", r)) //uncomment for cloud testing
}
