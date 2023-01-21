package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ex1() {
	fmt.Println("---ex1---")
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println("---ex1---")
}

func ex2() {
	fmt.Println("---ex2---")
	x := 10
	if x > 5 {
		fmt.Println(x)
		a, x := 20, 5
		fmt.Println(a, x)
	}
	fmt.Println(x)
	fmt.Println("---ex2---")
}

func ex3() {
	fmt.Println("---ex3---")

	rand.Seed(time.Now().UnixNano())

	if n := rand.Intn(10); n == 0 {
		fmt.Println(n)
		fmt.Println("n is 0")
	} else {
		fmt.Println(n)
		fmt.Println("n is not 0")
	}

	fmt.Println("---ex3---")
}

func main() {
	ex1()
	ex2()
	ex3()
}
