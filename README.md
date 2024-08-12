# Reusable

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/reusable.svg)](https://pkg.go.dev/github.com/akramarenkov/reusable)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/reusable)](https://goreportcard.com/report/github.com/akramarenkov/reusable)
[![codecov](https://codecov.io/gh/akramarenkov/reusable/branch/master/graph/badge.svg?token=)](https://codecov.io/gh/akramarenkov/reusable)

## Purpose

Library with reusable temporary buffer of variable length

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/reusable/grower"
    "github.com/akramarenkov/reusable/simple"
)

func main() {
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
```
