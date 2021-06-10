package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/pandulaDW/go-microservice-with-grpc/data"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func sendSuccessMessage(w io.Writer) error {
	type response struct {
		Message string `json:"message"`
	}
	message := response{Message: "success"}
	encoder := json.NewEncoder(w)
	return encoder.Encode(message)
}

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
	if r.Method == http.MethodPut {
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)
		fmt.Println(g, r.URL.Path)
		if len(g) != 1 {
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}
		if len(g[0]) != 1 {
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "id should be an integer", http.StatusBadRequest)
		}
		p.l.Println(id)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println(http.MethodGet + ": " + r.URL.String())
	rw.Header().Set("Content-Type", "application/json")
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

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	err = sendSuccessMessage(rw)
	if err != nil {
		p.l.Fatal(err)
	}
}
