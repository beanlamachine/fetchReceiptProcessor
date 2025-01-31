package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type receiptWithPointsandid struct {
	id           string `json:"id"`
	retailer     string `json:"retailer"`
	purchaseDate string `json:"purchaseDate"`
	purchaseTime string `json:"purchaseTime"`
	items        []item `json:"items"`
	total        string `json:"total"`
	points       int    `json:"points"`
}

type receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []item `json:"items"`
	Total        string `json:"total"`
}

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var receipts = []receipt{}

var processedReceipts = []receiptWithPointsandid{}

func processReceipts(c *gin.Context) {

	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	receipts = append(receipts, newReceipt)

	newReceiptWithID := receiptWithPointsandid{
		id:           uuid.New().String(),
		retailer:     newReceipt.Retailer,
		purchaseDate: newReceipt.PurchaseDate,
		purchaseTime: newReceipt.PurchaseTime,
		items:        newReceipt.Items,
		total:        newReceipt.Total,
		points:       calculatePoints(&newReceipt),
	}

	processedReceipts = append(processedReceipts, newReceiptWithID)

	c.JSON(http.StatusCreated, gin.H{"id": newReceiptWithID.id})
}

func getPointsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range processedReceipts {
		if a.id == id {
			c.IndentedJSON(http.StatusOK, a.points)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipts)
	router.GET("/receipts/:id/points", getPointsByID)

	router.Run("localhost:8080")
}

// Functions to calculate points
func calculatePoints(receipt *receipt) int {
	points := 0
	points += countAlphanumericCharacters(receipt.Retailer)
	points += checkRoundDollarAmount(receipt.Total)
	points += checkMultipleOfQuarter(receipt.Total)
	points += 5 * (len(receipt.Items) / 2)
	for _, itm := range receipt.Items {
		points += calculateItemPoints(itm)
	}
	points += checkOddDay(receipt.PurchaseDate)
	points += checkPurchaseTime(receipt.PurchaseTime)
	return points
}

func countAlphanumericCharacters(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func checkRoundDollarAmount(total string) int {
	var dollarAmount float64
	if _, err := fmt.Sscanf(total, "%f", &dollarAmount); err == nil {
		if dollarAmount == math.Floor(dollarAmount) {
			return 10
		}
	}
	return 0
}

func checkMultipleOfQuarter(total string) int {
	var dollarAmount float64
	if _, err := fmt.Sscanf(total, "%f", &dollarAmount); err == nil {
		if math.Mod(dollarAmount, 0.25) == 0 {
			return 25
		}
	}
	return 0
}

func calculateItemPoints(item item) int {
	trimmedDesc := strings.TrimSpace(item.ShortDescription)
	if len(trimmedDesc)%3 == 0 {
		var price float64
		fmt.Sscanf(item.Price, "%f", &price)
		adjustedPrice := math.Ceil(price * 0.2)
		return int(adjustedPrice)
	}
	return 0
}

func checkOddDay(purchaseDate string) int {
	parsedDate, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return 0
	}
	if parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

func checkPurchaseTime(purchaseTime string) int {
	parsedTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0
	}
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}
