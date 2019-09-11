package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/elazarl/goproxy"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Println("Listen 127.0.0.1:7070")
	log.Fatal(http.ListenAndServe("127.0.0.1:7070", proxy))
}
