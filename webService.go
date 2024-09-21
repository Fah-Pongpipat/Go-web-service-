package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ProductList []product

func intProduct() {
	products := `[
		{"id": 102156, "name": "Apple", "price": 10.0},
		{"id": 201879, "name": "Banana", "price": 5.0},
		{"id": 332445, "name": "Orange", "price": 7.0}
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
		if newProdict.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProdict.ID = checkID()
		ProductList = append(ProductList, newProdict)
		w.WriteHeader(http.StatusCreated)
		return

	}
}

func checkID() int {
	newID := 0
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Intn(999999)
	for _, product := range ProductList {
		if product.ID != randomID {
			newID = randomID
		}
	}
	return newID
}
func main() {
	intProduct()
	fmt.Print(ProductList)
	http.HandleFunc("/products", productHandler)
	http.ListenAndServe(":3000", nil)
}
