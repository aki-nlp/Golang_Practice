package main

import "fmt"

// 最初に呼ばれる関数
func init() {
	fmt.Println("Init!")
}

// 普通の関数
func f1() {
	fmt.Println("f1!!")
}

// 必ず必要な関数
func main() {
	fmt.Println("Hello World!")
	f1()
}

/* ----出力----
Init!
Hello World!
f1!!
   ----------- */
