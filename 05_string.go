package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello"

	// 表示
	fmt.Println(s)
	fmt.Println(s[0])         // Hが文字コードが表示される
	fmt.Println(string(s[0])) // 文字として表示
	fmt.Println(s[0:2])
	fmt.Println(s + s)
	fmt.Println(len(s), "\n")
	// fmt.Println(s[-1]) -1 はできない
	// fmt.Println(s * 2) できない

	// 文字コードを表示
	fmt.Println('a', "\n")

	// 全て小文字, 大文字で表示
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s), "\n")

	// replace
	fmt.Println(strings.Replace(s, "l", "x", 1))
	fmt.Println(strings.ReplaceAll(s, "l", "x"), "\n")

	// 含まれるか
	fmt.Println(strings.Contains(s, "x"), "\n")

	// そのまま表示
	fmt.Println(`H
 e
  l
   l
    o`)

}

/* ----出力----
Hello
72
H
He
HelloHello
5

97

hello
HELLO

Hexlo
Hexxo

false

H
 e
  l
   l
    o
   ----------- */
