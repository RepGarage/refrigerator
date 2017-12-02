package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/ngoncharov/refrigerator/db"
)

var (
	productsBaseURL     = os.Getenv("PRODUCTS_BASE_URL")
	productImageBaseURL = os.Getenv("PRODUCT_IMAGE_BASE_URL")
	partnerID           = os.Getenv("PARTNER_ID")
)

// ApiInterface api calls
type ApiInterface struct{}

// ProductsResponse type response
type ProductsResponse struct {
	Text    string
	Results []ResponseProduct
}

// ResponseProduct type product entity
type ResponseProduct struct {
	Entity    string `json:"entity"`
	Text      string `json:"text"`
	ProductID int    `json:"product_id,omitempty"`
}

// GetProducts return products object
func (p ApiInterface) GetProducts(name string) ([]db.Product, error) {
	r, err := http.Get(productsBaseURL + "prompt?text=" + name)
	if err != nil {
		return []db.Product{}, err
	}
	defer r.Body.Close()

	var productsResponse ProductsResponse
	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return []db.Product{}, e
	}

	_err := json.Unmarshal(body, &productsResponse)
	if _err != nil {
		return []db.Product{}, _err
	}

	result := convertProductsResponseToProducts(productsResponse)
	return result, nil
}

func convertProductsResponseToProducts(r ProductsResponse) []db.Product {
	var result []db.Product
	for _, v := range r.Results {
		if v.Entity == "product" {
			result = append(
				result,
				db.Product{
					Name:      v.Text,
					ProductID: v.ProductID,
				},
			)
		}
	}
	return result
}

// GetProductImageFromAPI returns image of product
func (p ApiInterface) GetProductImageFromAPI(side string, productID interface{}) (string, error) {
	var id string
	switch v := productID.(type) {
	case int:
		result := strconv.Itoa(v)
		id = result
		break
	case string:
		id = v
	}
	r, e := http.Get(productImageBaseURL + "partner/" + partnerID + "/item/" + id + "/picture/?format=jpg&width=" + side + "&height=" + side + "&scale=both&encoding=binary")
	if e != nil {
		return "", e
	}
	defer r.Body.Close()
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	resultImageFormat := "data:image/jpeg;base64, " + b64.StdEncoding.EncodeToString(response)
	return resultImageFormat, nil
}
