package chap10

import "fmt"

func goroutine() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch := make(chan int, len(a))
	for _, v :=range a {
		// 引数に渡す
		go func(val int) {
			ch <- val* 2
		}(v)
	}

	for i :=0; i < len(a); i ++ {
		fmt.Println(<-ch, " ")
	}
}

func Main() {
	goroutine()
}