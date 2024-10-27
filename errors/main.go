package main

import (
	"errors"
	"fmt"
	"log"
)

func calculateArea(length, width int) (int, error) {
	if length <= 0 || width <= 0 {
		return 0, errors.New("length and width must be positive")
	}
	return length * width, nil
}

var ErrInsufficientFunds = fmt.Errorf("insufficient funds")
var ErrInvalidQuantity = fmt.Errorf("invalid quantity")

func purchaseItem(price, quantity, balance int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	} else if price*quantity > balance {
		return ErrInsufficientFunds
	}
	return nil
}

func main() {
	length, width := 5, 10
	area, err := calculateArea(length, width)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Area of rectangle with length %d and width %d is %d\n", length, width, area)

	price, quantity, balance := 100, 2, 500
	err = purchaseItem(price, quantity, balance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Purchase of %d items at $%d each successful. Remaining balance: $%d\n", quantity, price, balance-quantity*price)
}
