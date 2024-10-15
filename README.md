# Reusable

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/reusable.svg)](https://pkg.go.dev/github.com/akramarenkov/reusable)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/reusable)](https://goreportcard.com/report/github.com/akramarenkov/reusable)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/reusable/badge.svg)](https://coveralls.io/github/akramarenkov/reusable)

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
