package Checker

import (
	"errors"
	"fmt"
)

func Check(args string) error {
	if len(args) == 0 {
		// fmt.Println("EMPTY ARGS")
		return errors.New("Empty args")
	}
	for index, v := range args {
		if !(v <= 126) {
			fmt.Printf("INVALID ARGS in index %d\n", index)
			return errors.New("Invalid input")
		}
	}
	return nil
}

func ChekcOnlyN(args string) (bool, int) {
	counter := 0
	flag := true
	// if len(args) != 1 {

	for i := 0; i < len(args)-1; i++ {
		if args[i] == '\\' && args[i+1] == 'n' {
			counter += 2
		}
		if args[i] != '\\' && args[i] != 'n' {
			flag = false
		}
	}

	if counter == len(args) && flag {
		return false, counter / 2
	} else {
		return true, counter
	}
} // else {
//	return true, counter
//}
//}
