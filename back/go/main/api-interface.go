package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/RepGarage/refrigerator/back/go/main/db"
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

type httpClient interface {
	Get(string) (*http.Response, error)
}

// GetProducts return products object
func (p ApiInterface) GetProducts(httpClient httpClient, name string) ([]db.Product, error) {
	r, err := httpClient.Get(productsBaseURL + "prompt?text=" + name)
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
func (p ApiInterface) GetProductImageFromAPI(httpClient httpClient, side interface{}, productID interface{}) (string, error) {
	var id string
	var sSide string
	// Converting
	switch v := productID.(type) {
	case int:
		id = strconv.Itoa(v)
		break
	case string:
		id = v
	}
	switch v := side.(type) {
	case string:
		sSide = v
		break
	case int:
		sSide = strconv.Itoa(v)
		break
	}

	r, e := httpClient.Get(productImageBaseURL + "partner/" + partnerID + "/item/" + id + "/picture/?format=jpg&width=" + sSide + "&height=" + sSide + "&scale=both&encoding=binary")
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
