# Reusable

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/reusable.svg)](https://pkg.go.dev/github.com/akramarenkov/reusable)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/reusable)](https://goreportcard.com/report/github.com/akramarenkov/reusable)
[![codecov](https://codecov.io/gh/akramarenkov/reusable/branch/master/graph/badge.svg?token=h4cv2z4hnB)](https://codecov.io/gh/akramarenkov/reusable)

## Purpose

Library with reusable temporary buffer of variable length

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/reusable"
)

func main() {
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
```
