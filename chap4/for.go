package chap4

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for1()
	for2()
	// for3()
	// for4()
	for5()
	for6()
	for7()

	ex1()
	ex2()
	ex3()

	switch1()
	println("----")
	switch2()
}

func for1() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func for2() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}

func for3() {
	for {
		println("forever")
	}
}

func for4() {
	for {
		i := 0
		if i == 5 {
			break
		}
		i += 1
	}
}

func for5() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		println(i)
	}
}

func for6() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		println(i, v)
	}

	// Mapの場合は順不同
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	for _, v := range ages {
		println(v)
	}

	str := "hello"
	for i, v := range str {
		println(i, v)
		println(i, string(v))
	}

	kanji := "これは日本語です"
	for i, v := range kanji {
		println(i, v, string(v))
	}
}

func for7() {
	samples := []string{"a", "b", "c"}
outer:
	for _, s := range samples {
		for _, r := range s {
			if r == 'b' {
				continue outer
			}
			println(string(r))
		}
	}
}



func ex1() {
	fmt.Println("---ex1---")
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println("---ex1---")
}

func ex2() {
	fmt.Println("---ex2---")
	x := 10
	if x > 5 {
		fmt.Println(x)
		a, x := 20, 5
		fmt.Println(a, x)
	}
	fmt.Println(x)
	fmt.Println("---ex2---")
}

func ex3() {
	fmt.Println("---ex3---")

	rand.Seed(time.Now().UnixNano())

	if n := rand.Intn(10); n == 0 {
		fmt.Println(n)
		fmt.Println("n is 0")
	} else {
		fmt.Println(n)
		fmt.Println("n is not 0")
	}

	fmt.Println("---ex3---")
}


func switch1() {
	intList := []int{1, 2, 3, 4, 5}
	for _, v := range intList {
		switch v {
		case 1:
			println("one")
		case 2:
			println("two")
		case 3:
			println("three")
		case 4:
			println("four")
		case 5:
			println("five")
		default:
			println("unknown")
		}
	}
}

func switch2() {
	intList := []int{1, 2, 3, 4, 5}
	for _, v := range intList {
		switch {
		case v == 1:
			println("one")
		case v == 2:
			println("two")
		case v == 3:
			println("three")
		case v == 4:
			println("four")
		case v == 5:
			println("five")
		default:
			println("unknown")
		}
	}
}
