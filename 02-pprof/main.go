package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func slowFunction(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintln(w, "OK")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", slowFunction)
	router.HandleFunc("/debug/pprof/", pprof.Index)
	addr := "127.0.0.1:8081"
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
