package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}

func B() {
	fmt.Println("b fonksiyonu çalıştı")
	defer A()
	defer C()
}
func C() {
	fmt.Println("c fonksiyonu çalıştı")
}
