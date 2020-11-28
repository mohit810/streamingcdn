package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
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
	fmt.Println("Server is Up and Running at Port:" + strconv.Itoa(*port))
	wd, _ := os.Getwd()
	fmt.Println(wd)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	if err != nil {
		panic(err)
	}
	//panic(http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert.pem", "key.pem", r))
}
