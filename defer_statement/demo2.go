package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi :", p.productName)
}

func Demo2() {
	p := product{productName: "Laptop", unitPrice: 100}
	defer p.save()
	fmt.Println("İşlem başarılı")
}
