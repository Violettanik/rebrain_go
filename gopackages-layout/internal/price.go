package internal

import "errors"

const DEFAULT_DISCOUNT = 30

type Customer struct {
    Name         string
    Age          int
    Balance      int
    Debt         int
    Discount     bool
    CalcDiscount func() (int, error)
}

// CalcPrice calculates the final price with discount
func CalcPrice(customer Customer, price int) (int, error) {
    if customer.CalcDiscount == nil {
        return price, nil
    }

    discount, err := customer.CalcDiscount()
    if err != nil {
        return 0, err
    }

    finalPrice := price - (price * discount / 100)
    if finalPrice < 0 {
        finalPrice = 0
    }

    return finalPrice, nil
}

// NewCustomer creates a new Customer with default CalcDiscount function
func NewCustomer(name string, age, balance, debt int, discount bool) Customer {
    cust := Customer{
        Name:     name,
        Age:      age,
        Balance:  balance,
        Debt:     debt,
        Discount: discount,
    }

    cust.CalcDiscount = func() (int, error) {
        if !cust.Discount {
            return 0, errors.New("discount not available")
        }
        result := DEFAULT_DISCOUNT - cust.Debt
        if result < 0 {
            return 0, nil
        }
        return result, nil
    }

    return cust
}
