package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3, 4} {
		fmt.Println(v)
		if v == 1 {
			fmt.Println("v = 1")
		} else if v == 2 {
			fmt.Println("v = 2")
		} else {
			fmt.Println("other number")
		}
	}

	fmt.Println("===")

	for _, v := range []int{1, 2, 3, 4} {
		fmt.Println(v)
		switch v {
		case 1:
			fmt.Println("v = 1")
			fallthrough
		case 2:
			fmt.Println("v = 2")
		default:
			fmt.Println("other number")
		}
	}

	fmt.Println("===")

	// おみくじ
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1
	fmt.Println("サイコロの目は", s, "です")
	switch s {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}

	fmt.Println("===")

	for i := 1; i <= 100; i = i + 1 {
		fmt.Print(i)
		fmt.Print("-")
		if i%2 == 0 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	}
}
