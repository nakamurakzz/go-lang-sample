package chap15

import (
	"fmt"
)

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

type Stack[T comparable] struct {
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

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

// アルゴリズムの抽象化
func Map[T1, T2 any] (s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 any] (s []T1, f func(T2, T1) T2, init T2) T2 {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any] (s []T, f func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func consoleLog[T any](val T) {
	fmt.Println(val)
}

func Main(){
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	// s.Push("hello") // これはコンパイルエラーになる

	for n := 0; n < 4; n++ {
		if v, ok := s.Pop(); ok {
			fmt.Println(v)
		} else {
			fmt.Println("stack is empty")
		}
	}

	consoleLog(1)
	consoleLog("hello")
	consoleLog(1.2)

	words := []string{"hello", "world", "foo", "bar"}
	fmt.Println(words)
	filtered := Filter(words, func(s string) bool {
		return s != "foo"
	})
	fmt.Println(filtered)

	words2 := []int{1, 2, 3, 4, 5}
	fmt.Println(words2)
	mapped := Map(words2, func(s int) int {
		return s * 10
	})
	fmt.Println(mapped)
	mapped2 := Reduce(words2, func(acc int, s int) int {
		return acc + s
	}, 0)
	fmt.Println(mapped2)

	// これはコンパイルエラーになる
}

// 型ターム
type BuildInType interface {
	string | int | float64
}

