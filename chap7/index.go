package chap7

import (
	"fmt"
	"time"
)

type MailCategory int

// constの()内の最初の行に型を定義して、値にiotaを指定する
// iotaは定数を0からインクリメントした値に紐づける
// 値に意味がある場合にはiotaを使うべきではない
// ただし、最初は0なので、0を意味する値がない場合は'_'を指定する
const (
	Uncategorized MailCategory = iota
	Personal
	Work
	Spam
	Social
	Advertising
)

type Status int

const (
	_ Status = iota
	Active
	Inactive
)

// IS-AよりもHAS-Aの方が良い
// (IS-Aはgoではできない)
type Employee struct {
	Id   int
	Name string
}

func (e Employee) Description() string {
	return fmt.Sprintf("ID: %d, Name: %s", e.Id, e.Name)
}

type Manager struct {
	Employee            // 変数名がつけられていないので埋め込みField
	Reports  []Employee // 部下
}

func (m Manager) FindNewEmployee() []Employee {
	newEmployee := []Employee{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
	}
	return newEmployee
}

func main() {
	m := Manager{
		Employee: Employee{Id: 3, Name: "Boss"},
		Reports:  []Employee{},
	}

	fmt.Println(m.Description())
	fmt.Println(m.Id)
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)

	m.Reports = m.FindNewEmployee()
	fmt.Println(m.Reports)
	fmt.Println(m.Employee.Id)

	method()
}

// 埋め込まれた側は、上位のメソッドを呼び出すことができない
// 継承ではなく、Composition的

type Stringer interface {
	String() string
}

type LogicProvider struct{}

type Person struct {
	FirsName string
	LastName string
	Age      int
}

// レシーバーを指定することでメソッドを定義できる
func (p Person) FullName() string {
	return p.FirsName + " " + p.LastName
}

func method() {
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
