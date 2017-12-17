package db

// Product model
type Product struct {
	ProductID int    `json:"_id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	URL       string `json:"url" bson:"url"`
}

//Photo model
type Photo struct {
	ProductID int    `json:"_id" bson:"_id"`
	Side      int    `json:"side" bson:"side"`
	Data      string `json:"data" bson:"data"`
}
