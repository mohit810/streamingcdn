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

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	r := httprouter.New()
	signal.HTTPSDPServer(r)
	fmt.Println("Server is Up and Running at Port:" + strconv.Itoa(*port))
	wd, _ := os.Getwd()
	fmt.Println(wd)
	panic(http.ListenAndServe(":"+strconv.Itoa(*port), r))
}
