package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	bob := person{name: "Bob", age: 20}
	fmt.Println(bob.name, bob.age)

	tom := person{}
	fmt.Println(tom)

	// 無名構造体
	p := struct {
		name string
		age  int
		pet  string
	}{
		name: "Bob",
		age:  20,
		pet:  "cat",
	}

	fmt.Println(p)

}
