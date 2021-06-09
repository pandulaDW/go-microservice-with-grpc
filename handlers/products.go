package handlers

import (
	"encoding/json"
	"github.com/pandulaDW/go-microservice-with-grpc/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetProducts()
	content, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write(content)
	if err != nil {
		p.l.Fatal(err)
	}
}
