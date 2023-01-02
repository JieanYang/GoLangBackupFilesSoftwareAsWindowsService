package main

import (
	"fmt"
	"net/http"
	"sync"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}
func startHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		http.HandleFunc("/", rootHandler)

		err := http.ListenAndServe(":9001", nil)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	wg.Wait()
}

func main () {
	startHttp()
}