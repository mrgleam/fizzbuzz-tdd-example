package fizzbuzz

import "strconv"

func FizzBuzz(number string) string {
	realNumber, _ := strconv.Atoi(number)
	if realNumber%3 == 0 {
		return "Fizz"
	}
	return number
}
