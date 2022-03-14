package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}

	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}
}
