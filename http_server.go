package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
	log.Println("/healthz Client IP Address is " + r.RemoteAddr + " http Status Code is " + http.StatusText(200))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		w.Header().Add(key, strings.Join(value, " "))
	}
	log.Println("/ Client IP Address is " + r.RemoteAddr + " http Status Code is " + http.StatusText(200))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
