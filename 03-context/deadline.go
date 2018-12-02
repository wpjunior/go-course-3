package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func handler01(w http.ResponseWriter, r *http.Request) {
	deadline := time.Now().Add(time.Second * 2)
	newCtx, cancel := context.WithDeadline(r.Context(), deadline)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		req, _ := http.NewRequest("GET", "http://localhost:8082/", nil)
		req = req.WithContext(newCtx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Ih deu erro1: ", err.Error())
			return
		}
		defer resp.Body.Close()
	}()

	go func() {
		defer wg.Done()

		req, _ := http.NewRequest("GET", "http://localhost:8082/", nil)
		req = req.WithContext(newCtx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Ih deu erro2: ", err.Error())
			return
		}
		defer resp.Body.Close()
	}()

	wg.Wait()
	fmt.Fprintln(w, "2 requests concluidos")
}

func handler02(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintln(w, "Request concluiidoo")
}

func server01() {
	addr := "127.0.0.1:8081"
	log.Printf("Running web server01 on: http://%s\n", addr)
	http.ListenAndServe(addr, http.HandlerFunc(handler01))
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
