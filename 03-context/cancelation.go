package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler01(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://localhost:8082/", nil)
	req = req.WithContext(r.Context())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func server01() {
	addr := "127.0.0.1:8081"
	log.Printf("Running web server01 on: http://%s\n", addr)
	http.ListenAndServe(addr, http.HandlerFunc(handler01))
}

func handler02(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/octet-stream")
	flusher := w.(http.Flusher)
	timeout := time.NewTimer(time.Second * 10)
	count := 0
	for {
		select {
		case <-r.Context().Done():
			fmt.Println("Request Cancelado")
			return

		case <-time.After(time.Second / 10):
			fmt.Fprintf(w, "[tsuru like] Please wait ........................................ %d/? \n\r", count)
			flusher.Flush()
			count++

		case <-timeout.C:
			fmt.Fprintln(w, "Request concluiidoo")
			return
		}
	}

}

func server02() {
	addr := "127.0.0.1:8082"
	log.Printf("Running proxy server02 on: http://%s\n", addr)
	http.ListenAndServe(addr, http.HandlerFunc(handler02))
}

func main() {
	go server01()
	go server02()
	select {}
}
