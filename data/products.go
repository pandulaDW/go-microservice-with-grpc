package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productList
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func AddProduct(product *Product) {
	product.ID = getNextId()
	productList = append(productList, product)
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return 0, ErrProductNotFound
}

func UpdateProduct(id int, product *Product) error {
	productIndex, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = id
	productList[productIndex] = product
	return nil
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky Coffee",
		Price:       2.45,
		SKU:         "abc345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd245",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
