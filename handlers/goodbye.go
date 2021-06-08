package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l: l}
}

func (gb *GoodBye) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	gb.l.Println("Good, bye!!")
	_, err := rw.Write([]byte("Bye bye"))
	if err != nil {
		gb.l.Fatal(err)
	}
}
