package captcha

import (
	"fmt"
)

var (
	operator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}

	operand = map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
)

func Generate(f, lo, op, ro int) string {
	result := map[int]string{
		0: fmt.Sprintf("%s %s %d", operand[lo], operator[op], ro),
		1: fmt.Sprintf("%d %s %s", lo, operator[op], operand[ro]),
	}
	return result[f]
}
