package main

import (
	cardvalidation "card-validator/card-validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type card struct {
	CardNumber string `json:"cardnumber"`
}

func main() {
	// Configure the api router
	router := gin.Default()

	// Configure Status path
	router.GET("/status", getStatus)

	// Configure validation path
	router.PUT("/validate", putCardNumber)

	// Run server
	router.Run("localhost:8080")
}

// Status path to check the status of the api
func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Operating Normally")
}

// validate credit card number using custom function
func putCardNumber(c *gin.Context) {
	var newCard card

	// Call BindJSON to bind the received JSON to newCard
	if err := c.BindJSON(&newCard); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid Request")
		return
	}

	// Validate card number using custom function
	isValid := cardvalidation.ValidateCardNumber(newCard.CardNumber)

	// Send custom response based on whether or not the card number is valid
	if isValid {
		c.IndentedJSON(http.StatusOK, "Card number is valid")
	} else {
		c.IndentedJSON(http.StatusOK, "Card number is invalid")
	}

}
