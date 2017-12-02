package main

import (
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

	result, err := apiInterface.GetProducts(requestProductName)

	assert.Equal(t, nil, err, "error should be nil")

	assert.Equal(t, 1, len(result), "result len should be 1")

}

func TestGerProductImageFromAPI(t *testing.T) {
	var testProductID = "1"
	var testProductSide = "100"
	defer gock.Off()
	gock.New(productImageBaseURL).
		Get("partner/" + partnerID + "/item/" + testProductID + "/picture/").
		Reply(200).
		BodyString("test")

	result, err := apiInterface.GetProductImageFromAPI(testProductSide, testProductID)

	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, true, len(result) > 0, "result length should be more than 0")
}
