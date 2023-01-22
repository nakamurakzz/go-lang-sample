package main

import "sort"

type Person struct {
	FirsName string
	LastName string
	Age      int
}

func main() {
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
