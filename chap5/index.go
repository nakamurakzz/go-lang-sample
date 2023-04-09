package chap5

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

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

	closure()

	deferF()

	function()

	funcVal()
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


type Person struct {
	FirsName string
	LastName string
	Age      int
}

func closure() {
	people := []Person{
		{"John", "Doe", 25},
		{"Jane", "Alen", 23},
		{"John", "Smith", 30},
		{"Jane", "Zwee", 28},
	}

	println("初期データ")
	for _, p := range people {
		println(p.FirsName, p.LastName, p.Age)
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	println("LastNameでソート")
	for _, p := range people {
		println(p.FirsName, p.LastName, p.Age)
	}

	// 年齢でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	println("Ageでソート")
	for _, p := range people {
		println(p.FirsName, p.LastName, p.Age)
	}

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i <= 5; i++ {
		println(i, twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(i int) int {
		return base * i
	}
}


func deferF() {
	if len(os.Args) < 2 {
		log.Fatal("引数が足りません")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
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

func function() {
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



func funcVal() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "/", "three"},
		{"2", "+", "three"},
		{"5", "a", "3"},
		{"5"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			println("1. invalid expression")
			continue
		}
		pl, err := strconv.Atoi(exp[0])
		onFunc, ok := opMap[exp[1]]
		if !ok {
			println("2. invalid operator")
			continue
		}
		pr, err := strconv.Atoi(exp[2])
		if err != nil {
			println("3. invalid expression")
			continue
		}
		result := onFunc(pl, pr)
		println(result)
	}
}

func add(i, j int) int { return 1 + j }
func sub(i, j int) int { return 1 - j }
func mul(i, j int) int { return 1 * j }
func div(i, j int) int { return 1 / j }

type onFuncType func(int, int) int

var opMap = map[string]onFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
