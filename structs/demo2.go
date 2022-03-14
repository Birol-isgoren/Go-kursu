package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Birol", lastName: "İşgören", age: 35}
	c.save()
}
