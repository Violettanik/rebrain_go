package main

import (
    "errors"
    "fmt"
    "gopackages-layout/internal"
)

func startTransactionDynamic(debtor interface{}) error {
    // Проверяем в рантайме, реализует ли аргумент интерфейс Debtor
    d, ok := debtor.(internal.Debtor)
    if !ok {
        return errors.New("incorrect type")
    }
    
    // Вызываем оригинальную логику списания долга
    return d.WrOffDebt()
}

func main() {
    // Создаем клиента с балансом 1000 и долгом 200
    customer := internal.NewCustomer("John", 1000, 200)
    
    // Корректный вызов (Customer реализует Debtor)
    err := startTransactionDynamic(customer)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println("Debt successfully written off")
    }
    
    // Вызов с неправильным типом
    wrongType := "I'm not a Debtor"
    err = startTransactionDynamic(wrongType)
    if err != nil {
        fmt.Printf("Error: %v\n", err) // Выведет "Error: incorrect type"
    }
}
