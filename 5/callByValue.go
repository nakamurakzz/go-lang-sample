package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	p := person{}
	i := 2
	s := "hello"
	fmt.Println(i, s, p)
	modFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{1: "one11", 2: "two22"}
	fmt.Println(m)
	modMap(m)
	fmt.Println(m)

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	modSlice(slice)
	fmt.Println(slice)
}

func modFails(i int, s string, p person) {
	i = 1
	s = "hello"
	p.name = "John"
}

func modMap(m map[int]string) {
	m[1] = "one"
	m[2] = "two"
	delete(m, 1)
}

func modSlice(s []int) {
	s[0] = 10
	s[1] = 20
	s[2] = 30
}
