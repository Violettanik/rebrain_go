package main

import (
	"fmt"
	"gopackages-layout/internal"
)

func main() {
	// Создаем клиента
	customer := internal.NewCustomer("John", 30, 1000, 200)
	
	// Проверяем реализацию интерфейса
	var debtor internal.Debtor = customer
	fmt.Printf("Type: %T\n", debtor) // *internal.Customer
	
	// Тестируем списание долга
	err := debtor.WrOffDebt()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Debt written off successfully")
	}
	
	// Попытка доступа к полям balance и debt напрямую:
	// Это вызовет ошибку компиляции:
	// fmt.Println(customer.balance) // ошибка - поле недоступно
	// fmt.Println(customer.debt)    // ошибка - поле недоступно
	
	// Доступ только к открытым полям:
	fmt.Println("Name:", customer.Name)
	fmt.Println("Age:", customer.Age)
}
