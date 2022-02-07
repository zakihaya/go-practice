package main

import "fmt"

func main() {
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum := n1 + n2 + n3 + n4
	fmt.Println(sum)

	fmt.Println("===========")

	numbers := []int{19, 86, 1, 12}

	total := 0
	for i := 0; i < len(numbers); i += 1 {
		total += numbers[i]
	}

	fmt.Println(total)

	fmt.Println("===========")

	total2 := 0
	for _, number := range numbers {
		total2 += number
	}

	fmt.Println(total2)
}
