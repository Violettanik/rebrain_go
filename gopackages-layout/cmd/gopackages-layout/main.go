package main

import (
	"fmt"
	"gopackages-layout/internal"
)

func main() {
	// Создаем клиента со скидкой
	customerWithDiscount := internal.NewCustomer("John", 30, 1000, 10, true)
	
	// Создаем клиента без скидки
	customerWithoutDiscount := internal.NewCustomer("Alice", 25, 500, 5, false)
	
	// Тестируем CalcPrice
	price := 1000
	
	// Клиент со скидкой (передаётся как Discounter)
	finalPrice, err := internal.CalcPrice(customerWithDiscount, price)
	if err != nil {
		fmt.Printf("Error for %s: %v\n", customerWithDiscount.Name, err)
	} else {
		fmt.Printf("Final price for %s: %d\n", customerWithDiscount.Name, finalPrice)
	}
	
	// Клиент без скидки (передаётся как Discounter)
	finalPrice, err = internal.CalcPrice(customerWithoutDiscount, price)
	if err != nil {
		fmt.Printf("Error for %s: %v\n", customerWithoutDiscount.Name, err)
	} else {
		fmt.Printf("Final price for %s: %d\n", customerWithoutDiscount.Name, finalPrice)
	}
}
