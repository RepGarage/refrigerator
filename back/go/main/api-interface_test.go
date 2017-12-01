package main_test

import (
	"encoding/json"
	"os"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

var (
	productsBaseURL     = os.Getenv("PRODUCTS_BASE_URL")
	productImageBaseURL = os.Getenv("PRODUCT_IMAGE_BASE_URL")
	partnerID           = os.Getenv("PARTNER_ID")
)

func TestGetProducts(t *testing.T) {
	var requestProductName = "сыр"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var mockResponse, _ = json.Marshal(
		struct {
			Text    string
			Results []struct {
				Entity    string `json:"entity"`
				Text      string `json:"text"`
				ProductID int    `json:"product_id,omitempty"`
			}
		}{
			Text: "",
			Results: {
				{
					Entity:    "product",
					Text:      "сыр",
					ProductID: 1,
				},
			},
		},
	)
	var jsonResponder, _ = httpmock.NewJsonResponder(200, mockResponse)

	httpmock.RegisterResponder(
		"GET",
		productsBaseURL+"prompt?text="+requestProductName,
		httpmock.NewJsonResponder(200, ``),
	)
}
