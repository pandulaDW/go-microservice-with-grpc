package handlers

import (
	"encoding/json"
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
		re := regexp.MustCompile(`/products/([0-9]+)$`)
		matches := re.FindSubmatch([]byte(r.URL.String()))

		var productId int
		if len(matches) > 0 {
			productId, _ = strconv.Atoi(string(matches[1]))
		} else {
			http.Error(rw, "invalid uri. product id must be provided", http.StatusBadRequest)
			return
		}
		p.updateProduct(productId, rw, r)
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

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println(http.MethodPut + ": " + r.URL.String())
	prod := new(data.Product)
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, data.ErrProductNotFound.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "unknown server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	err = sendSuccessMessage(rw)
	if err != nil {
		p.l.Fatal(err)
	}
}
