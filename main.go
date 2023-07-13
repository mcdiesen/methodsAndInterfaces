package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// printDetails Method of type Product
func (product *Product) printDetails() {
	fmt.Println("Name", product.name, "Category", product.category, "Price", product.price)
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (Supplier *Supplier) printDetails() {
	fmt.Println("Name", Supplier.name, "City", Supplier.city)
}
func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		p.printDetails()
	}
	Suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	fmt.Println("=====Suppliers====")
	for _, s := range Suppliers {
		s.printDetails()
	}
}
