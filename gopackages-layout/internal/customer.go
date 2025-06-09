package internal

import "errors"

type Debtor interface {
    WrOffDebt() error
}

// Overduer содержит финансовые поля
type Overduer struct {
    balance int
    debt    int
}

// Customer теперь содержит Overduer как встроенную структуру
type Customer struct {
    Name string
    Age  int
    Overduer
}

// WrOffDebt реализует интерфейс Debtor для Customer
func (c *Customer) WrOffDebt() error {
    if c.debt >= c.balance { // Обращение к полям через встроенную структуру
        return errors.New("not possible write off")
    }

    c.balance -= c.debt
    c.debt = 0

    return nil
}

// NewCustomer создает нового Customer
func NewCustomer(name string, age, balance, debt int) *Customer {
    return &Customer{
        Name: name,
        Age:  age,
        Overduer: Overduer{
            balance: balance,
            debt:    debt,
        },
    }
}
