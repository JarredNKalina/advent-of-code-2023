package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IsDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func ReadFileIntoString(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func SplitLines(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
}

func Min(values ...int) int {
	if len(values) == 0 {
		log.Fatal("min: empty slice")
	}

	minValue := values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Sum(elems []int) int {
	sum := 0

	for _, elem := range elems {
		sum += elem
	}

	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func FindLCM(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]

	for _, num := range numbers[1:] {
		result = lcm(result, num)
	}
	return result
}
