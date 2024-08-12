package simple_test

import (
	"fmt"

	"github.com/akramarenkov/reusable/grower"
	"github.com/akramarenkov/reusable/simple"
)

func ExampleBuffer() {
	buffer := simple.New[byte](1024, grower.Exactly)

	fmt.Println(len(buffer.Get(0)), cap(buffer.Get(0)))
	fmt.Println(len(buffer.Get(2048)), cap(buffer.Get(2048)))
	fmt.Println(len(buffer.Get(1024)), cap(buffer.Get(1024)))
	fmt.Println(len(buffer.Get(2048)), cap(buffer.Get(2048)))

	buffer.Reset()

	fmt.Println(len(buffer.Get(0)), cap(buffer.Get(0)))

	// Output:
	// 1024 1024
	// 2048 2048
	// 1024 2048
	// 2048 2048
	// 1024 1024
}
