package main

import (
	cardvalidation "card-validator/card-validation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type card struct {
	CardNumber string `json:"cardnumber"`
}

func main() {
	fmt.Print(cardvalidation.ValidateCardNumber("5122301102209636"))
	router := gin.Default()

	router.GET("/status", getStatus)

	router.Run("localhost:8080")
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Operating Normally")
}
}
