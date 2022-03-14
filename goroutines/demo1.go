package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		time.Sleep(time.Millisecond)
		fmt.Println("Çift sayı : ", i)
	}
}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		time.Sleep(time.Millisecond)
		fmt.Println("Tek sayı : ", i)
	}
}
