package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d -- %s", b.parameter, b.message)
}

func TahmitEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: tahmin, message: "Sınırlar dışında bir değer girdiniz"}
	}
	return "Bildiniz", nil
}
