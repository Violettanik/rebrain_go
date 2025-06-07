package internal

import "errors"

const DEFAULT_DISCOUNT = 30

// Discounter интерфейс для типов, которые могут рассчитывать скидку
type Discounter interface {
    CalcDiscount() (int, error)
}

// Customer реализует интерфейс Discounter
type Customer struct {
    Name     string
    Age      int
    Balance  int
    Debt     int
    Discount bool
}

// CalcDiscount реализует метод интерфейса Discounter
func (c Customer) CalcDiscount() (int, error) {
    if !c.Discount {
        return 0, errors.New("discount not available")
    }
    result := DEFAULT_DISCOUNT - c.Debt
    if result < 0 {
        return 0, nil
    }
    return result, nil
}

// CalcPrice рассчитывает итоговую цену с учетом скидки
// Теперь принимает любой объект, реализующий интерфейс Discounter
func CalcPrice(d Discounter, price int) (int, error) {
    discount, err := d.CalcDiscount()
    if err != nil {
        return 0, err
    }

    finalPrice := price - (price*discount)/100
    if finalPrice < 0 {
        finalPrice = 0
    }

    return finalPrice, nil
}

// NewCustomer создает нового Customer (фабричная функция)
func NewCustomer(name string, age, balance, debt int, discount bool) Customer {
    return Customer{
        Name:     name,
        Age:      age,
        Balance:  balance,
        Debt:     debt,
        Discount: discount,
    }
}
