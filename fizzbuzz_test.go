package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		testcase, number, expected string
	}{
		{"input 1 should 1", "1", "1"},
	}
	for _, test := range tests {
		if result := FizzBuzz(test.number); result != test.expected {
			t.Fatalf("%s FizzBuzz(%s) = %v, want %v", test.testcase, test.number, result, test.expected)
		}
	}
}
