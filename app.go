package main

import (
	cardvalidation "card-validator/card-validation"
	"fmt"
)

type card struct {
	CardNumber string `json:"cardnumber"`
}

func main() {
	fmt.Print(cardvalidation.ValidateCardNumber("5122301102209636"))
}
