package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello, World!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = fmt.Fprintf(rw, "Received message: %s", d)
	if err != nil {
		h.l.Fatal(err)
	}
}
