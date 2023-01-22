package main

import "fmt"

func main() {
	// ポインタも変数
	// その中身が別の変数が保存されているアドレス
	var x int32 = 10
	var y bool = true

	// 変数x,yが使用しているメモリのアドレスを表示
	pointerX := &x
	pointerY := &y
	var pointerZ *string

	// &xはxのアドレスを表す
	// ポインタ型
	pointerToX := &x

	// *は関節参照演算子
	// ポインタが指す先の値を取得する
	fmt.Println(x, y, pointerX, pointerY, pointerZ, pointerToX, *pointerToX)

	p := 20
	var pointerToP *int
	pointerToP = &p
	fmt.Println(p, pointerToP, *pointerToP)

	var f *int
	failedUpdate(f)
	fmt.Println(f)

	x2 := 10
	fmt.Println(&x2)
	failedUpdate(&x2)
	fmt.Println(x2)
	update(&x2)
	fmt.Println(x2)
}

func failedUpdate(g *int) { // gはここでコピーされる
	x := 1
	g = &x
}

func update(g *int) {
	fmt.Println(&*g)
	*g = 1 // *gはgのコピー元のメモリに格納されている値を表す
}

type Foos struct {
	bar string
}
