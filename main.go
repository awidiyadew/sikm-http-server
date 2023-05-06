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

func GetProduct() ([]model.Product, error) {
	file, err := os.OpenFile("data/product.txt", os.O_RDONLY, 0644)
	if err != nil {
		return []model.Product{}, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []model.Product{}, err
	}
	var products []model.Product
	spiltter := strings.Split(string(content), "\n")
	for _, RowContact := range spiltter {
		dats := strings.Split(RowContact, "_")
		prc, _ := strconv.ParseFloat(dats[2], 64)
		qts, _ := strconv.Atoi(dats[3])
		product := model.Product{
			ID:    dats[0],
			Name:  dats[1],
			Price: prc,
			Qty:   qts,
		}
		products = append(products, product)
	}
	return products, nil

}
func GetProductList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		products, _ := GetProduct()
		jsonResp, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)

	}
}

func AddProduct(c model.Product) (model.Product, error) {
	file, err := os.OpenFile("data/product.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return model.Product{}, err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + c.ID + "_" + c.Name + "_" + strconv.FormatFloat(c.Price, 'f', 2, 64) + "_" + strconv.Itoa(c.Qty))
	if err != nil {
		return model.Product{}, err
	}
	return c, nil

}

func AddProductList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var c model.Product

		if c.Qty <= 0 {
			message := "qty must be greater than 0"
			messageJson, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageJson)
			return 
		}

		err = json.Unmarshal(reqBody, &c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct, err := AddProduct(c)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, _ := json.Marshal(newProduct)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
func main() {
	http.HandleFunc("/product", GetProductList())
	http.HandleFunc("/product/add", AddProductList())
	http.ListenAndServe("localhost:3000", nil)

}
