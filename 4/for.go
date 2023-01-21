package main

func main() {
	for1()
	for2()
	// for3()
	// for4()
	for5()
	for6()
	for7()
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
