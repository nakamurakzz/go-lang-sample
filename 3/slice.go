package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s[1:3])
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	var s1 [][]int
	s1 = append(s1, []int{1, 2, 3}) // appendは非推奨

	fmt.Println(s1)
	fmt.Println(len(s1))

	var s2 []int
	if s2 == nil {
		fmt.Println("s2 is nil")
	}

	var s3 = []int{}
	if s3 == nil {
		fmt.Println("s3 is nil")
	}

	var ss = [][]int{s, s}
	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}

	// フルスライス
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x) // [1 2 3 4 60]
	fmt.Println("y:", y) // [1 2 30 40 50]
	fmt.Println("z:", z) // [3 4 70]

	x1 := []int{1, 2, 3, 4, 5}
	y1 := make([]int, 4) // 4つの要素を持つスライスを作成
	num := copy(y1, x1)
	x1[0] = 100
	fmt.Println(x1)
	fmt.Println(num, y1) // 4 [1 2 3 4]

	// Map
	tmap := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(tmap["apple"])
	delete((tmap), "apple")
	fmt.Println(tmap["apple"])

	// Like Set
	intSet := map[int]bool{} // 空のマップを作成、ゼロ値はfalse
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet[11])
}
