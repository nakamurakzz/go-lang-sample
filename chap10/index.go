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

// 文字列sに関して、searchersに含まれる関数を並列に実行し、最も早く終了した結果を返す
// 第一引数は対象の文字列
// 第二引数は文字列を受け取って文字列のスライスを返す関数のスライス
// 返り値は文字列のスライス
func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	resultChan := make(chan []string)
	for _, searcher := range searchers {
		go func(f func(string) []string){
			select {
			case resultChan <- f(s):
				fmt.Println("resultChan")
			case <-done:
				fmt.Println("done")
			}
		}(searcher)
	}

	r := <-resultChan
	close(done)
	return r
}

func Main() {
	goroutine()
}