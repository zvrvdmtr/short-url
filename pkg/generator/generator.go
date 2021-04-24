package generator

import (
	"strings"
	"math"
)

var mapping = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ShortUrlGenerator(id int) string {
	char_array := strings.Split(mapping, "")
	numArray := bijectiveFunction(len(char_array), id)
	result := make([]string, 0)
	for i := len(numArray)-1; i>=0; i-- {
		result = append(result, char_array[numArray[i]])
	}
	return strings.Join(result, "")
}

func bijectiveFunction(length int, id int) []int{
	result := make([]int, 0)
	for id > 0 {
		reminder := id % length
		result = append(result, reminder)
		id = id / length
	}
	return result
}

func BijectiveDecode(short string) int {
	result := 0
	for i:=0; i < len(short); i++ {
		result += strings.Index(mapping, string(short[i])) * int(math.Pow(float64(len(mapping)), float64(len(short)-1-i)))
	}
	return result
}