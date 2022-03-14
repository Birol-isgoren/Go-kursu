package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "Birol"
	fmt.Println(strings.Count(isim, "i"))
	fmt.Println(strings.Contains(isim, "o"))
	fmt.Println(strings.Index(isim, "i"))
}
