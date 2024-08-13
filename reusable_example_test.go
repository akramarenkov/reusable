package reusable_test

import (
	"fmt"

	"github.com/akramarenkov/reusable"
	"github.com/akramarenkov/reusable/grower"
)

func ExampleBuffer_limited() {
	buffer := reusable.New[byte](1024)

	for _, length := range []int{1024, 1000, 2048} {
		slice := buffer.Get(length)
		fmt.Println(len(slice), cap(slice))
	}

	// Output:
	// 1024 1024
	// 1000 1024
	// 1024 1024
}

func ExampleBuffer_unlimited() {
	buffer := reusable.New[byte](0)

	for _, length := range []int{1024, 1472, 2048, 1024} {
		slice := buffer.Get(length)
		fmt.Println(len(slice), cap(slice))
	}

	// Output:
	// 1024 1472
	// 1472 1472
	// 2048 2752
	// 1024 2752
}

func ExampleBuffer_custom_growing() {
	buffer := reusable.New[byte](0, grower.Quarter)

	for _, length := range []int{1024, 1280, 2048, 1024} {
		slice := buffer.Get(length)
		fmt.Println(len(slice), cap(slice))
	}

	// Output:
	// 1024 1280
	// 1280 1280
	// 2048 2560
	// 1024 2560
}
