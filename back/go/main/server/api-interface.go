package server

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/centrypoint/refrigerator/back/go/main/db"
)

var (
	productsBaseURL     = os.Getenv("PRODUCTS_BASE_URL")
	productImageBaseURL = os.Getenv("PRODUCT_IMAGE_BASE_URL")
	partnerID           = os.Getenv("PARTNER_ID")
)

// APIInterface api calls
type APIInterface interface {
	GetProducts(httpClient HTTPClient, name string) ([]db.Product, error)
	GetProductImageFromAPI(httpClient HTTPClient, side interface{}, productID interface{}) (string, error)
	GetProductShelfLife(httpClient HTTPClient, url string) (string, error)
}

// API ApiInterface implementation
type API struct{}

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
	URL       string `json:"url,omitempty"`
}

// HTTPClient interface
type HTTPClient interface {
	Get(string) (*http.Response, error)
}

// GetProducts return products object
func (p API) GetProducts(httpClient HTTPClient, name string) ([]db.Product, error) {
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
					URL:       v.URL,
				},
			)
		}
	}
	return result
}

// GetProductImageFromAPI returns image of product
func (p API) GetProductImageFromAPI(httpClient HTTPClient, side interface{}, productID interface{}) (string, error) {
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

// GetProductShelfLife return product shelf life
func (p API) GetProductShelfLife(httpClient HTTPClient, url string) (string, error) {
	var err error
	var res string
	var response *http.Response
	var body []byte

	if response, err = httpClient.Get(productsBaseURL + url); err != nil {
		return "", err
	}

	defer response.Body.Close()
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return "", err
	}

	var bodyRows = strings.Split(string(body), "\n")
	log.Println("FOCKEN ERROR")

	for i, v := range bodyRows {

		// fmt.Println(v)
		if strings.HasPrefix(strings.Trim(v, " "), "<h4") && strings.Contains(v, "Условия хранения") {
			for index, value := range bodyRows[i : i+60] {
				if strings.Contains(value, "Срок хранения макс.") || strings.Contains(value, "Срок хранения") {
					// for _, valueI := range bodyRows[i+index : i+index+7] {
					// 	fmt.Println(strings.Trim(valueI, " "))
					// }
					if strings.HasPrefix(strings.Trim(bodyRows[i+index+3], " "), "<a") {
						res = strings.Trim(bodyRows[i+index+5], " ")
					} else {
						res = strings.Trim(bodyRows[i+index+4], " ")
					}
					break
				}
			}
			break
		}
	}

	return res, err
}
