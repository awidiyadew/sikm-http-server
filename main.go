package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"sikm-http-server/model"
	"sikm-http-server/service"
)

/**
1. Read File  // library Os
2. Construct data from file to struct
3. Construct struct to json // library json marshall
4. Create response // w.writer
**/

func GetProduct() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// handling error with proper method (GET ONLY) 
		if r.Method != "GET"{
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("The method you are using to access API is not allowed"))
			return
		 }
		data, _ := service.GetProduct()
		jsonData, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to marshal JSON"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

	}
}

/**
1. Get new data from request body (format JSON)
2. Construct JSON to struct // json unmarshall
3. Open file // library Os
4. Append new data to file // create new function to utilize this service
5. Construct new data from struct to JSON // json marshall
6. Create response // w.writer
**/
func PostProduct() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// handling error with proper method (POST ONLY)
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

		// Handling invalid request (qty must be above 0)
		if p.Quantity <= 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			errBody,_ := json.Marshal(map[string]string{
				"message": "qty must be more than 0",
			})
			w.Write(errBody)
			return
		}

		insertedProduct, err := service.AddProduct(p)
		if err != nil {
			// http.Error(w, "failed add contact", http.StatusInternalServerError)
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
