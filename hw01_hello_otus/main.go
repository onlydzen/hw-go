package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	data := "Hello, OTUS!"
	result := stringutil.Reverse(data)

	fmt.Println(result)
}
