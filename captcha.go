package captcha

import (
	"errors"
	"fmt"
)

var (
	operator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}

	operand = map[int]string{
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

func generateCaptcha(f, lo, op, ro int) (string, error) {
	opr, ok := operator[op]
	if !ok {
		return "", errors.New("unsupport operator")
	}

	if f == 0 {
		lor, ok := operand[lo]
		if !ok {
			return "", errors.New("unsupport left operand")
		}

		return fmt.Sprintf("%s %s %d", lor, opr, ro), nil
	}

	if f == 1 {
		ror, ok := operand[ro]
		if !ok {
			return "", errors.New("unsupport right operand")
		}

		return fmt.Sprintf("%d %s %s", lo, opr, ror), nil
	}

	return "", errors.New("unsupport format")
}
