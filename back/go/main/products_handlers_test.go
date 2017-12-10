package main_test

import (
	"net/http"
	"net/url"
	"testing"

	main "github.com/centrypoint/refrigerator/back/go/main"
	"github.com/centrypoint/refrigerator/back/go/main/db"
)

var (
	testProductsApi          main.ProductsAPI
	apiInterfaceMockInstance apiInterfaceMock
	mockHTTPWriterInstance   mockHTTPWriter
	dbInterfaceMockInstance  DBInterfaceMock
	testName200              = "200"
	testName500              = "500"
)

/**
* Mocks
 */
// Mock api interface
type apiInterfaceMock struct{}

func (m apiInterfaceMock) GetProducts(c main.HTTPClient, name string) ([]db.Product, error) {
	switch name {
	case testName200:
		return []db.Product{
			{
				Name:      "Product",
				ProductID: 1,
			},
		}, nil
	default:
		return nil, nil
	}
}

func (p apiInterfaceMock) GetProductImageFromAPI(httpClient main.HTTPClient, side interface{}, productID interface{}) (string, error) {
	return "", nil
}

// Mock http writer
type mockHTTPWriter struct{}

func (m mockHTTPWriter) Write([]byte) (int, error) {
	return 0, nil
}
func (m mockHTTPWriter) WriteHeader(int) {

}

func (m mockHTTPWriter) Header() http.Header {
	return map[string][]string{"test": {"header"}}
}

// Database interface mock
type DBInterfaceMock struct{}

func (d DBInterfaceMock) Connect(string) error {
	return nil
}

func (d DBInterfaceMock) InsertProductToDB(p db.Product) error {
	return nil
}

func (d DBInterfaceMock) FetchPhotoFromDBByProductID(interface{}, interface{}) (db.Photo, error) {
	return db.Photo{}, nil
}

func (d DBInterfaceMock) FetchProductFromDBByID(interface{}) (db.Product, error) {
	return db.Product{}, nil
}

func (d DBInterfaceMock) InsertPhotoToDB(photo db.Photo) error {
	return nil
}

/**
* Tests
 */

func TestGetProductsHandler(t *testing.T) {
	var mockURL, _ = url.Parse("http://test.com/test?name=" + testName200)
	var mockRequest = http.Request{
		URL: mockURL,
	}
	testProductsApi.GetProductsHandler(apiInterfaceMockInstance, dbInterfaceMockInstance, mockHTTPWriterInstance, &mockRequest)
}
