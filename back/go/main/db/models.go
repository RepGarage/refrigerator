package db

// Product model
type Product struct {
	ProductID int    `json:"product_id" bson:"product_id"`
	Name      string `json:"name" bson:"name"`
}

//Photo model
type Photo struct {
	ProductID int    `json:"product_id" bson:"product_id"`
	Side      int    `json:"side" bson:"side"`
	Data      string `json:"data" bson:"data"`
}
