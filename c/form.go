package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
}
