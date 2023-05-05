package main

import (
	"sikm-http-server/model"
	"io/ioutil"
	"encoding/json"
	"sikm-http-server/service"
	"net/http"
)

func GetProduct()  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := service.GetProduct()
		jsonResp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}

func PostProduct()  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)
			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		var p model.Product
		err = json.Unmarshal(reqBody, &p)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		insertedProduct, err := service.AddProduct(p)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errBody, _ := json.Marshal(map[string]string{
				"message": err.Error(),
			})
			w.Write(errBody)
			return
		}
		resp, _ := json.Marshal(insertedProduct)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func main() {
	http.HandleFunc("/product/list", GetProduct())
	http.HandleFunc("/product/add", PostProduct())

	http.ListenAndServe(":3000", nil)
}