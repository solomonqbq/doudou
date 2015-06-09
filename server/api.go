package main

import (
	"flag"
	"fmt"
	"github.com/xeniumd-china/vitamin/network"
	"net/http"
)

var port string

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	ip := network.GetLocalIP()
	fmt.Fprintf(w, "Hello world!My Ip is %s and serving on port %s\n", ip, port)
}

func main() {
	port = "8080"
	flag.StringVar(&port, "p", "", "port to listen")
	flag.Parse()

	mux := &MyMux{}
	http.ListenAndServe(":"+port, mux)
}
