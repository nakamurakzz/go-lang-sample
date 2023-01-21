package main

func main() {
	switch1()
	println("----")
	switch2()
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
