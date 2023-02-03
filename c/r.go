package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Carot struct {
	num    int
	weight int
	sugar  int
}

func main() {
	// 標準入力（1行目）をスペース区切りで分割
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()
	slice := strings.Split(line, " ")

	n, _ := strconv.Atoi(slice[0]) // 人参の個数
	s, _ := strconv.Atoi(slice[1]) // 目安の糖分
	p, _ := strconv.Atoi(slice[2]) // 許容値

	answerCarot := Carot{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := sc.Text()
		slice := strings.Split(line, " ")

		weight, _ := strconv.Atoi(slice[0])
		sugar, _ := strconv.Atoi(slice[1])

		if sugar >= s-p && sugar <= s+p {
			if answerCarot.num == 0 {
				answerCarot.num = i + 1
				answerCarot.weight = weight
				answerCarot.sugar = sugar
				continue
			}

			if weight < answerCarot.weight {
				answerCarot.num = i + 1
				answerCarot.weight = weight
				answerCarot.sugar = sugar
			}
		}
	}
	if answerCarot.num == 0 {
		fmt.Printf("not found")
		return
	}
	fmt.Printf("%d", answerCarot.num)

}
