package internal

import "errors"

type Debtor interface {
    WrOffDebt() error
}

type Customer struct {
    name    string
    balance int
    debt    int
    // другие поля
}

func (c *Customer) WrOffDebt() error {
    if c.debt >= c.balance {
        return errors.New("not possible write off")
    }

    c.balance -= c.debt
    c.debt = 0

    return nil
}

func NewCustomer(name string, balance, debt int) *Customer {
    return &Customer{
        name:    name,
        balance: balance,
        debt:    debt,
    }
}
