package service

import (
	"fmt"
	"io/ioutil"
	"sikm-http-server/model"
	"strings"
	"os"
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
	arrOfData := strings.Split(string(content), "\n")
	for _, rowProduct := range arrOfData{
		data := strings.Split(rowProduct, "_")
		product := model.Product{
			ID:  data[0],
			Name:   data[1],
			Price: data[2],
			Quantity: data[3],
		}
		products = append(products, product)
	}

	return products, nil
}

func AddProduct(p model.Product) (model.Product, error) {
	file, err := os.OpenFile("data/product.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return model.Product{}, err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "\n%v_%v_%v_%v", p.ID, p.Name, p.Price, p.Quantity)
	if err != nil {
		return model.Product{}, err
	}

	return p, nil
}