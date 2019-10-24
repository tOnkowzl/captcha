package main

import (
	"fmt"
	"strconv"

	"github.com/divan/num2words"
)

func generateCaptcha(f, lo, op, ro int) string {
	lor := strconv.Itoa(lo)
	ror := strconv.Itoa(ro)

	switch f {
	case 0:
		lor = num2words.Convert(lo)
	case 1:
		ror = num2words.Convert(ro)
	}

	opr := "+"
	switch op {
	case 0:
		opr = "+"
	case 1:
		opr = "-"
	case 2:
		opr = "*"
	}

	return fmt.Sprintf("%s %s %s", lor, opr, ror)
}
