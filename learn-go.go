package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":1818", "http service address") // Q=17, R=18
var count int

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(hello))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	// templ.Execute(w, req.FormValue("s"))
	log.SetPrefix(time.Now().Format("2017-03-09 15:51:57.89"))
	log.Print("Starting handling request.....")
	log.Print("RemoteAddr: ", req.RemoteAddr)
	log.Print("RequestURI: ", req.RequestURI)
	log.Print("Host: ", req.Host)
	log.Print("Method: ", req.Method)
	log.Print("Header", req.Header)

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}

	log.Print("Body: ", string(bodyBytes))

	count++
	fmt.Fprint(w, "hello ", count)
	log.Print("Handler exists.")
}
