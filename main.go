package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type receipt struct {
	id           string `json:"id"`
	retailer     string `json:"retailer"`
	purchaseDate string `json:"purchaseDate"`
	purchaseTime string `json:"purchaseTime"`
	items        []item `json:"items"`
	total        string `json:"total"`
	points       string `json:"points"`
}

type item struct {
	shortDescription string `json:"shortDescription"`
	price            string `json:"price"`
}

var receipts = []receipt{
	{
		id:           "1",
		retailer:     "Target",
		purchaseDate: "2022-01-01",
		purchaseTime: "13:01",
		items: []item{
			{shortDescription: "Mountain Dew 12PK", price: "6.49"},
			{shortDescription: "Emils Cheese Pizza", price: "12.25"},
			{shortDescription: "Knorr Creamy Chicken", price: "1.26"},
			{shortDescription: "Doritos Nacho Cheese", price: "3.35"},
			{shortDescription: "Klarbrunn 12-PK 12 FL OZ", price: "12.00"},
		},
		total:  "35.35",
		points: "20",
	},
}

func processReceipts(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	newReceipt.id = uuid.New().String()

	receipts = append(receipts, newReceipt)

	c.JSON(http.StatusCreated, gin.H{"id": newReceipt.id})

}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipts)

	router.Run("localhost:8080")
}
