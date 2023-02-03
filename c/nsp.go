package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 標準入力（1行目）をスペース区切りで分割
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()
	slice := strings.Split(line, " ")

	n, _ := strconv.Atoi(slice[0]) // 全体の件数
	s, _ := strconv.Atoi(slice[1]) // ページあたりの件数
	p, _ := strconv.Atoi(slice[2]) // 目的のページ数

	pages := make([][]int, n/s+1)
	pageNum := 0
	for i := 1; i <= n; i++ {
		pages[pageNum] = append(pages[pageNum], i)
		if i%s == 0 {
			pageNum++
		}
	}

	if len(pages[p-1]) == 0 {
		fmt.Println("none")
		return
	}

	for i, page := range pages[p-1] {
		if i == len(pages[p-1])-1 {
			fmt.Print(page)
		} else {
			fmt.Print(page, " ")
		}
	}
}
