package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		testcase, number, expected string
	}{
		{"input 1 should 1", "1", "1"},
		{"input 3 should Fizz", "3", "Fizz"},
		{"input 6 should Fizz", "6", "Fizz"},
		{"input 9 should Fizz", "9", "Fizz"},
		{"input 5 should Buzz", "5", "Buzz"},
	}
	for _, test := range tests {
		if result := FizzBuzz(test.number); result != test.expected {
			t.Fatalf("%s FizzBuzz(%s) = %v, want %v", test.testcase, test.number, result, test.expected)
		}
	}
}
