package main

import (
	"domain-as-logic/domain"
	"fmt"
)

func main() {
	// Создаем регион
	moscow := domain.Region{
		Code:         "MSK",
		TaxRate:      0.20, // НДС 20%
		ShippingRate: 300,  // 300 рублей за кг
	}

	// Создаем корзину
	cart := domain.NewCart(moscow)

	// Создаем товары
	laptop := domain.Product{
		ID:       "laptop-1",
		Name:     "MacBook Pro",
		Price:    150000,
		Weight:   2.0,
		Category: "electronics",
	}

	phone := domain.Product{
		ID:       "phone-1",
		Name:     "iPhone 15",
		Price:    90000,
		Weight:   0.5,
		Category: "electronics",
	}

	// Добавляем товары в корзину
	cart.AddItem(laptop, 1)
	cart.AddItem(phone, 2)

	// Добавляем скидки
	cart.AddDiscount(domain.Discount{
		Type:      "percentage",
		Value:     10,            // 10% скидка
		Category:  "electronics", // только на электронику
		MinAmount: 200000,        // при покупке от 200000
	})

	cart.AddDiscount(domain.Discount{
		Type:      "fixed",
		Value:     5000,   // фиксированная скидка 5000
		MinAmount: 100000, // при покупке от 100000
	})

	// Вычисляем и выводим детали
	subtotal := cart.CalculateSubtotal()
	discounts := cart.CalculateDiscounts(subtotal)
	shipping := cart.CalculateShipping()
	afterDiscounts := subtotal - discounts
	tax := cart.CalculateTax(afterDiscounts + shipping)
	total := cart.CalculateTotal()

	fmt.Printf("Детали заказа:\n")
	fmt.Printf("Промежуточная сумма: %.2f руб.\n", subtotal)
	fmt.Printf("Скидки: %.2f руб.\n", discounts)
	fmt.Printf("Доставка: %.2f руб.\n", shipping)
	fmt.Printf("НДС: %.2f руб.\n", tax)
	fmt.Printf("Итоговая сумма: %.2f руб.\n", total)
}
