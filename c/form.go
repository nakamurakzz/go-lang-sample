package c

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Salarry struct {
	amount int
	tax    int
}

func (s *Salarry) setTax() {
	// s.amountをfloatに変換
	amount := float64(s.amount)
	switch {
	case s.amount <= 100000:
		s.tax = 0
		return
	case s.amount <= 750000:
		s.tax = int((amount - 100000) * 0.1)
		return
	case s.amount <= 1500000:
		s.tax = int(65000 + (amount-750000)*0.2)
		return
	default:
		s.tax = int(215000 + (amount-1500000)*0.4)
		return
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num := sc.Text()

	// numをintに変換
	n, _ := strconv.Atoi(num)

	srallyArray := make([]Salarry, n)
	totalTax := 0
	for _, s := range srallyArray {
		sc.Scan()
		amount := sc.Text()
		amountInt, _ := strconv.Atoi(amount)
		s.amount = amountInt
		s.setTax()
		totalTax += s.tax
	}
	fmt.Println(totalTax)

	nsp()

	r()
}

func nsp () {
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

type Carot struct {
	num    int
	weight int
	sugar  int
}

func r() {
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
