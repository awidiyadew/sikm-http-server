package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"sikm-http-server/model"
	"strconv"
	"strings"
)

func GetProduct()([]model.Product, error) {
	file, err := os.OpenFile("data/product.txt", os.O_RDONLY, 0644)
	if err != nil {
		return []model.Product{}, err
	}

	defer file.Close()

	rawData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return []model.Product{}, err2
	}

	// when there's no data in the file return empty array of struct
	if len(rawData) == 0{
		return []model.Product{}, nil
	}
	var products []model.Product
	arrOfData := strings.Split(string(rawData), "\n") // 0001_Kit Kat_3500_10
	for _, rowProduct := range arrOfData {
		data := strings.Split(rowProduct, "_") // 0001 Kit Kat 3500 10 // 0 1 2 3 ...
		quantity,_ := strconv.Atoi(data[3])
		product := model.Product{
			ID: data[0],
			Name: data[1],
			Price: data[2],
			Quantity: quantity,
		}
		products = append(products, product)
	}

	return products, nil
}

func AddProduct(p model.Product) (model.Product, error){
	file, err := os.OpenFile("data/product.txt", os.O_APPEND|os.O_WRONLY, 0644) // to add new data in the last index
	if err != nil {
		return model.Product{}, err
	}
	
	defer file.Close()

	_, err = fmt.Fprintf(file, "\n%v_%v_%v_%v", p.ID, p.Name, p.Price, p.Quantity) // 0004_Sari Roti_4000_24
	if err != nil {
		return model.Product{}, err
	}

	return p, nil
}