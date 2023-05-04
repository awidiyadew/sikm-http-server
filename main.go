package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sikm-http-server/model"
	"strconv"
	"strings"
)

func GetProductList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}

		file, err := os.OpenFile("data/product.txt", os.O_RDONLY, 0644)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}
		var products []model.Product
		splittedProducts := strings.Split(string(content), "\n")
		for _, product := range splittedProducts {
			if product != "" {
				splittedProduct := strings.Split(product, "_")
				price, _ := strconv.ParseFloat(splittedProduct[2], 64)
				qty, _ := strconv.Atoi(splittedProduct[3])
				product := model.Product{
					Id:    splittedProduct[0],
					Name:  splittedProduct[1],
					Price: price,
					Qty:   qty,
				}
				products = append(products, product)
			}
		}
		productsJson, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productsJson)
	}
}

func AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		var product model.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "failed to decode JSON", http.StatusInternalServerError)
			return
		}
		if product.Qty <= 0 {
			message := "qty must be greater than 0"
			messageJson, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageJson)
			return
		}

		file, err := os.OpenFile("data/product.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		_, err = file.WriteString(product.Id + "_" + product.Name + "_" + strconv.FormatFloat(product.Price, 'f', 2, 64) + "_" + strconv.Itoa(product.Qty) + "\n")
		if err != nil {
			http.Error(w, "failed to write to file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product added successfully"))
	}
}

func main() {

	http.HandleFunc("/product/list", GetProductList())
	http.HandleFunc("/product/add", AddProduct())

	http.ListenAndServe(":8080", nil)

}
