package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Product struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
    Qty   int    `json:"qty"`
}

func productListHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    data, err := ioutil.ReadFile("data/product.txt")
    if err != nil {
        http.Error(w, "Failed to read data", http.StatusInternalServerError)
        return
    }

    var products []Product
    if err := json.Unmarshal(data, &products); err != nil {
        http.Error(w, "Failed to unmarshal data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}

func productAddHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var product Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    if product.Qty <= 0 {
        http.Error(w, "qty must be more than 0", http.StatusBadRequest)
        return
    }

    data, err := ioutil.ReadFile("data/product.txt")
    if err != nil {
        http.Error(w, "Failed to read data", http.StatusInternalServerError)
        return
    }

    var products []Product
    if err := json.Unmarshal(data, &products); err != nil {
        http.Error(w, "Failed to unmarshal data", http.StatusInternalServerError)
        return
    }

    products = append(products, product)

    newData, err := json.Marshal(products)
    if err != nil {
        http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
        return
    }

    if err := ioutil.WriteFile("data/product.txt", newData, 0644); err != nil {
        http.Error(w, "Failed to write data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

func main() {
    http.HandleFunc("/product/list", productListHandler)
    http.HandleFunc("/product/add", productAddHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}
