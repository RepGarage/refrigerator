package server_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	server "github.com/centrypoint/refrigerator/back/go/main/server"

	"github.com/stretchr/testify/assert"

	"gopkg.in/h2non/gock.v1"
)

var (
	productsBaseURL     = os.Getenv("PRODUCTS_BASE_URL")
	productImageBaseURL = os.Getenv("PRODUCT_IMAGE_BASE_URL")
	partnerID           = os.Getenv("PARTNER_ID")
	api                 server.API
)

func TestGetProducts(t *testing.T) {
	var requestProductName = "cheese"
	defer gock.Off()
	gock.New(productsBaseURL).
		Get("prompt").
		Reply(200).
		JSON(server.ProductsResponse{
			Text: "Test",
			Results: []server.ResponseProduct{
				{
					Entity:    "product",
					Text:      requestProductName,
					ProductID: 1,
				},
			},
		})

	result, err := api.GetProducts(&http.Client{}, requestProductName)

	assert.Equal(t, nil, err, "error should be nil")

	assert.Equal(t, 1, len(result), "result len should be 1")

}

type mockHTTPClient struct{}

func (m mockHTTPClient) Get(url string) (*http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
	}

	return response, nil
}

func TestGerProductImageFromAPI(t *testing.T) {
	var testProductID = "1"
	var testProductSide = "100"
	var mockHTTP mockHTTPClient
	gock.New(productImageBaseURL).
		Get("partner/" + partnerID + "/item/" + testProductID + "/picture/").
		Reply(200).
		BodyString("test")

	// Side and ProductID string
	resultString, err := api.GetProductImageFromAPI(&mockHTTP, testProductSide, testProductID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, len(resultString) > 0, "result length should be more than 0")

	// Side and ProductID int
	resultInt, err := api.GetProductImageFromAPI(&http.Client{}, 100, 1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, len(resultInt) > 0, "result length should be more than 0")
	gock.Off()
}


type shelfLifeHTTPClient struct {}
func (s shelfLifeHTTPClient) Get(url string) (*http.Response, error) {
	switch url {
	case "/moloko-syr-yaytsa/syry/hochl-syr-pl-tost-s-vet-lomt-45150g--305146?searchPhrase=ветчина":
		return http.Response{
			Body: 
		}
	}
}
func TestGetProductShelfLife(t *testing.T) {
	var err error
	var result string
	var api server.API
	var cases = map[string]string{
		"/moloko-syr-yaytsa/syry/hochl-syr-pl-tost-s-vet-lomt-45150g--305146?searchPhrase=ветчина":                   "180 дней",
		"/moloko-syr-yaytsa/syry/hochl-syr-55-s-vetch-vann-plavl-400g--308011?searchPhrase=ветчина":                  "180 дней",
		"/konservy-orehi-sousy/myasnye-konservy/elinskiy-vetchina-sterilizov-gost-325g--313010?searchPhrase=ветчина": "365 дней",
		"/catalog/moloko-syr-yaytsa/moloko/prav-mol-moloko-past-3-2-4-pet-2l--310201":                                "25 дней",
	}

	for k, v := range cases {
		if result, err = api.GetProductShelfLife(http.Client{}, k); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, v, result)
	}
}
