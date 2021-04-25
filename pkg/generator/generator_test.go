package generator

import (
	"fmt"
	"reflect"
	"testing"
)

type bijectiveTestData struct {
	length, id int
	result []int
	url string
}

var testData = []bijectiveTestData{
	{62, 12345, []int{7, 13, 3}, "dnh"},
	{62, 125, []int{1, 2}, "cb"},
	{62, 19158, []int{0, 61, 4}, "e9a"},
}

func TestPositiveShortUrlGenerator(t *testing.T) {
	for _, test := range testData {
		output := ShortUrlGenerator(test.id)
		if output != test.url {
			t.Errorf("got %v, wanted %v", output, test.url)
		}
	}
}

func TestPositiveBijectiveFunction(t *testing.T) {
	for _, test := range testData {
		output := bijectiveFunction(test.length, test.id)
		if !reflect.DeepEqual(output, test.result) {
			t.Errorf("got %v, wanted %v", output, test.result)
		}
	}
}

func TestPositiveBijectiveDecode(t *testing.T) {
	for _, test := range testData {
		output := BijectiveDecode(test.url)
		if test.id != output {
			t.Errorf("got %v, wanted %v", output, test.id)
		}
	}
}

func Benchmark(b *testing.B) {
	for i:=0; i<b.N; i++ {
		bijectiveFunction(62, 125)
	}
}


func ExampleShortUrlGenerator() {
	fmt.Println(ShortUrlGenerator(125))
	// Output: cb
}


func ExamplebijectiveFunction() {
	fmt.Println(bijectiveFunction(62, 125))
	// Output: []int{1, 2}
}


func ExampleBijectiveDecode() {
	fmt.Println(BijectiveDecode("cb"))
	// Output: 125
}