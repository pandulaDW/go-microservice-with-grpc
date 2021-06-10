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
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println(http.MethodGet + ": " + r.URL.String())
	rw.Header().Add("Content-Type", "application/json")
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println(http.MethodPost + ": " + r.URL.String())
	newProduct := new(data.Product)
	err := newProduct.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to decode the content", http.StatusBadRequest)
	}
	data.AddProduct(newProduct)

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write([]byte(`{ "message": "success"}`))
	if err != nil {
		p.l.Fatal(err)
	}
}
