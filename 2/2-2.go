package main

import "fmt"

func main() {
	const str = "Hello, 世界"

	fmt.Println(str)

	const (
		a = iota
		b
		c
	)

	const (
		d = 1 << iota
		e
	)

	fmt.Println(a, b, c, d, e)
}
