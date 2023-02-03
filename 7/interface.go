package main

// インターフェース
// - インターフェースも型の一種
// インターフェースの命名規則：名詞＋er

type Stringer interface {
	String() string
}

type LogicProvider struct{}
