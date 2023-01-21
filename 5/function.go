package main

import (
	"errors"
	"fmt"
	"os"
)

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type MyFuncOps struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOps) bool {
	print(opts.Age)
	return true
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	println(div(10, 2))
	println(div(10, 0))
	println("----")

	me := MyFuncOps{
		FirstName: "John",
		LastName:  "Doe",
		Age:       20,
	}
	MyFunc(me)
	println("----")
	a := []int{1, 2, 3, 4, 5}
	println(addTo(10, a...))

	println("----")
	// res, rem, err := divAndRemainder(10, 3)
	res, rem, err := divAndRemainder(10, 2)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	println(res, rem)
	println("----")
	res1, rem1, err1 := divAndRemainderReturnWithName(10, 3)
	if err1 != nil {
		println(err1)
		os.Exit(1)
	}
	println(res1, rem1)
}

/*
* 戻り値にエラーを返す
 */
func divAndRemainder(n, d int) (int, int, error) {
	if d == 0 {
		return 0, 0, fmt.Errorf("divided by zero")
	}
	return n / d, n % d, nil
}

func divAndRemainderReturnWithName(n, d int) (res int, rem int, err error) {
	if d == 0 {
		return n, d, errors.New("divided by zero")
	}
	res, rem = n/d, n%d
	return res, rem, nil
}
