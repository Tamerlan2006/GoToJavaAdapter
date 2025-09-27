package main

/*
#include <stdbool.h>
*/
import "C"

import (
	"fmt"
	"os"
	"strconv"
)

//export Validate
func Validate(input *C.char) C.bool {
	str := C.GoString(input)
	var sum int
	for i := len(str) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid character in input:")
			os.Exit(1)
		}
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + 1
			}
		}
		sum += digit
	}
	return C.bool(sum%10 == 0)
}

func main() {}
