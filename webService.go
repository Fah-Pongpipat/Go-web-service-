package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ProductList []product

func intProduct() {
	products := `[
		{"id": 1, "name": "Apple", "price": 10.0},
		{"id": 2, "name": "Banana", "price": 5.0},
		{"id": 3, "name": "Orange", "price": 7.0}
	]`

	err := json.Unmarshal([]byte(products), &ProductList)
	if err != nil {
		log.Fatal(err)
	}
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	prodcutjson, err := json.Marshal(ProductList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(prodcutjson)
	case http.MethodPost:
		var newProdict product
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &newProdict)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ProductList = append(ProductList, newProdict)
		w.WriteHeader(http.StatusCreated)
		return

	}
}
func main() {
	intProduct()
	fmt.Print(ProductList)
	http.HandleFunc("/products", productHandler)
	http.ListenAndServe(":3000", nil)
}
