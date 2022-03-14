package string_functions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Birol"
	fmt.Println(strings.HasPrefix(isim, "Bir"))
	fmt.Println(strings.HasSuffix(isim, "el"))
	fmt.Println(strings.Index(isim, "c"))
	harfler := []string{"b", "i", "r"}
	sonuc := strings.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(strings.Replace(sonuc, "*", "-", -1))
	fmt.Println(strings.Split(sonuc, "*"))
	fmt.Println(strings.Repeat(sonuc, 5))
}
