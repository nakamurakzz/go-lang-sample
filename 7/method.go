package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirsName string
	LastName string
	Age      int
}

// レシーバーを指定することでメソッドを定義できる
func (p Person) FullName() string {
	return p.FirsName + " " + p.LastName
}

func main() {
	p := Person{"John", "Doe", 25}
	println(p.FullName())

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	a := Adder{start: 1}
	fmt.Println(a.Add(2))

	// メソッド値
	am := a.Add
	fmt.Println(am(3))

	// メソッド式
	af := Adder.Add
	fmt.Println(af(a, 4))

}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %s", c.total, c.lastUpdated)
}

type Adder struct {
	start int
}

func (a Adder) Add(b int) int {
	return a.start + b
}
