package domain

// Product представляет товар в магазине
type Product struct {
	ID       string
	Name     string
	Price    float64
	Weight   float64
	Category string
}

// CartItem представляет товар в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Discount представляет скидку
type Discount struct {
	Type      string // "percentage" или "fixed"
	Value     float64
	Category  string  // если скидка применяется к определенной категории
	MinAmount float64 // минимальная сумма для применения скидки
}

// Region представляет регион с его особенностями
type Region struct {
	Code         string
	TaxRate      float64
	ShippingRate float64 // стоимость доставки за кг
}

// Cart представляет корзину покупок
type Cart struct {
	Items     []CartItem
	Discounts []Discount
	Region    Region
}

// NewCart создает новую корзину
func NewCart(region Region) *Cart {
	return &Cart{
		Items:     make([]CartItem, 0),
		Discounts: make([]Discount, 0),
		Region:    region,
	}
}

// AddItem добавляет товар в корзину
func (c *Cart) AddItem(product Product, quantity int) {
	// Проверяем, есть ли уже такой товар в корзине
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += quantity
			return
		}
	}

	// Если товара нет, добавляем новый
	c.Items = append(c.Items, CartItem{
		Product:  product,
		Quantity: quantity,
	})
}

// AddDiscount добавляет скидку в корзину
func (c *Cart) AddDiscount(discount Discount) {
	c.Discounts = append(c.Discounts, discount)
}

// CalculateSubtotal вычисляет промежуточную сумму без скидок и налогов
func (c *Cart) CalculateSubtotal() float64 {
	var subtotal float64
	for _, item := range c.Items {
		subtotal += item.Product.Price * float64(item.Quantity)
	}
	return subtotal
}

// CalculateDiscounts вычисляет сумму скидок
func (c *Cart) CalculateDiscounts(subtotal float64) float64 {
	var totalDiscount float64

	for _, discount := range c.Discounts {
		if subtotal < discount.MinAmount {
			continue
		}

		switch discount.Type {
		case "percentage":
			// Если скидка процентная
			if discount.Category == "" {
				// Скидка на всю корзину
				totalDiscount += subtotal * (discount.Value / 100)
			} else {
				// Скидка на категорию
				var categorySum float64
				for _, item := range c.Items {
					if item.Product.Category == discount.Category {
						categorySum += item.Product.Price * float64(item.Quantity)
					}
				}
				totalDiscount += categorySum * (discount.Value / 100)
			}
		case "fixed":
			// Если скидка фиксированная
			totalDiscount += discount.Value
		}
	}

	return totalDiscount
}

// CalculateShipping вычисляет стоимость доставки
func (c *Cart) CalculateShipping() float64 {
	var totalWeight float64
	for _, item := range c.Items {
		totalWeight += item.Product.Weight * float64(item.Quantity)
	}
	return totalWeight * c.Region.ShippingRate
}

// CalculateTax вычисляет налог
func (c *Cart) CalculateTax(amount float64) float64 {
	return amount * c.Region.TaxRate
}

// CalculateTotal вычисляет итоговую сумму
func (c *Cart) CalculateTotal() float64 {
	// Вычисляем промежуточную сумму
	subtotal := c.CalculateSubtotal()

	// Вычисляем скидки
	discounts := c.CalculateDiscounts(subtotal)

	// Вычисляем сумму после скидок
	afterDiscounts := subtotal - discounts

	// Вычисляем доставку
	shipping := c.CalculateShipping()

	// Вычисляем налог на сумму после скидок + доставку
	tax := c.CalculateTax(afterDiscounts + shipping)

	// Итоговая сумма
	total := afterDiscounts + shipping + tax

	return total
}
