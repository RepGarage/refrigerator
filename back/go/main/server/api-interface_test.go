package server_test

import (
	"bytes"
	"encoding/json"
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

type getProductsTestHTTPClient struct{}

func (h getProductsTestHTTPClient) Get(string) (*http.Response, error) {
	var (
		err  error
		body []byte
	)

	if body, err = json.Marshal(server.ProductsResponse{
		Text: "Test",
		Results: []server.ResponseProduct{
			{
				Entity:    "product",
				Text:      requestProductName,
				ProductID: 1,
			},
		},
	}); err != nil {
		return
	}

	return nil, &http.Response{
		Body: body,
	}
}

func TestGetProducts(t *testing.T) {
	var (
		requestProductName = "cheese"
		result             *http.Response
		httpClient         getProductsTestHTTPClient
		err                error
	)

	result, err = api.GetProducts(&httpClient, requestProductName)

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
