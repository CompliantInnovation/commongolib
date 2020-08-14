package commongolib

import (
	"fmt"
	"strconv"
)

func AE2Decode(input int) (int, bool) {
	original := input / 100
	ok := AE2Encode(original) == input
	return original, ok
}

func AE2Encode(input int) int {
	sum, count := sumDigits(input)
	for sum > 9 {
		sum, _ = sumDigits(sum)
	}

	for count > 9 {
		_, count = sumDigits(count)
	}

	result, err := strconv.Atoi(fmt.Sprintf("%d%d%d", input, count, sum))

	if err != nil {
		panic(err)
	}

	return result
}

func sumDigits(number int) (int, int) {
	remainder := 0
	sumResult := 0
	digitCount := 0
	for number != 0 {
		digitCount++
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult, digitCount
}
