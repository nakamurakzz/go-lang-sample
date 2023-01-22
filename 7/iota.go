package main

import "fmt"

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
}

// 埋め込まれた側は、上位のメソッドを呼び出すことができない
// 継承ではなく、Composition的
