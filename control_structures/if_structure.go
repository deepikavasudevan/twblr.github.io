package control_structures

import "strconv"

func fizzBuzz(i int) string {
	if i%3 == 0 && i%5 != 0  {
		return "Fizz"
	} else if i%3 != 0 && i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	}

	return strconv.Itoa(i)
}
