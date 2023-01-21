package main

import "strconv"

func main() {
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
