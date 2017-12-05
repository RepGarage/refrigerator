package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/h2non/gock.v1"
)

func TestGetProducts(t *testing.T) {
	var requestProductName = "cheese"
	defer gock.Off()
	gock.New(productsBaseURL).
		Get("prompt").
		Reply(200).
		JSON(ProductsResponse{
			Text: "Test",
			Results: []ResponseProduct{
				{
					Entity:    "product",
					Text:      requestProductName,
					ProductID: 1,
				},
			},
		})

	result, err := apiInterface.GetProducts(&http.Client{}, requestProductName)

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
	resultString, err := apiInterface.GetProductImageFromAPI(&mockHTTP, testProductSide, testProductID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, len(resultString) > 0, "result length should be more than 0")

	// Side and ProductID int
	resultInt, err := apiInterface.GetProductImageFromAPI(&http.Client{}, 100, 1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, len(resultInt) > 0, "result length should be more than 0")
	gock.Off()
}
