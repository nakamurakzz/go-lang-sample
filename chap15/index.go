package chap15

import "fmt"

// ジェネリクス

func Min(v1, v2 float64) float64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

type Node struct {
	val int
	next *Node
}

type Orderable interface {
	// メソッドOrderは次の値を返す
	// Orderableがvより小さい場合は負の値
	// Orderableがvと等しい場合は0
	// Orderableがvより大きい場合は正の値
	Order(v Orderable) int
}

type Tree struct {
	val Orderable
	left, right *Tree
}

func (t *Tree) Insert(v Orderable) *Tree {
	if t == nil {
		return &Tree{val: v}
	}

	switch comp := v.Order(t.val); {
		case comp < 0:
			t.left = t.left.Insert(v)
		case comp > 0:
			t.right = t.right.Insert(v)
		}
		return t
}

type OrderableInt int

// interface{} はJSのanyと同じ、型システムを活かせない
func (oi OrderableInt) Order(v interface{}) int {
	return int(oi - v.(OrderableInt))
}

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(v T) {
	s.vals = append(s.vals, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	v := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return v, true
}

func Main(){
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for n := 0; n < 4; n++ {
		if v, ok := s.Pop(); ok {
			fmt.Println(v)
		} else {
			fmt.Println("stack is empty")
		}
	}
}