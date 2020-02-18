package fmtprac

import "fmt"

func FmtPrac() {
	fmt.Printf("数値:%v, 文字列:%v, 配列:%v", 2, "test", [...]int{1, 2, 3})    // 複数の型を受け取ることができる
	fmt.Printf("数値:%#v, 文字列:%#v, 配列:%#v", 2, "test", [...]int{1, 2, 3}) // リテラルで表示
	fmt.Printf("数値:%T, 文字列:%T, 配列:%T", 2, "test", [...]int{1, 2, 3})    // 型を表示する
}
