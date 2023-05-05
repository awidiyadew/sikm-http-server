package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"sikm-http-server/model"
	"strings"
)

func GetProductList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 (Method Not Allowed)"))
			return
		}

		product, err := os.OpenFile("data/product.txt", os.O_RDONLY, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed Open Data Product"))
			return
		}
		defer product.Close()

		var prdk []model.Product
		scan := bufio.NewScanner(product)
		for scan.Scan() {
			line := scan.Text()
			fields := strings.Split(line, "_")
			if len(fields) != 4 {
				continue
			}

			list_produk := model.Product{
				Id: fields[0],
				Name: fields[1],
				Price: fields[2],
				Qty: fields[3],
			}

			prdk = append(prdk, list_produk)
		}
		
		response, _ := json.Marshal(prdk)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

type SuccessResponse struct {
	Message  string `json:"message"`
}

func AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 (Method Not Allowed)"))
			return
		}

		var product model.Product

		dataProduct, err := os.OpenFile("data/product.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to open users data file"))
			return
		}
		defer dataProduct.Close()

		_, err = dataProduct.WriteString("\n" + product.Id + "_" + product.Name + "_" + product.Price + "_" + product.Qty)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to add product"))
			return
		}

		success := SuccessResponse{
			Message:  "add user success",
		}
		response, err := json.Marshal(success)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Internal Server Error"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func main() {
	http.HandleFunc("/product/list", GetProductList())
	http.HandleFunc("/product/add", AddProduct())

	http.ListenAndServe(":8080", nil)
}