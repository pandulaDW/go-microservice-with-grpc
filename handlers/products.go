package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(rw)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
