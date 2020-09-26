package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 変数(var)宣言の仕方
	var i1 int
	i1 = 1
	var i2 int = 1
	var i3 = 1
	i4 := 1 // 初期値
	fmt.Println(i1, i2, i3, i4)
	fmt.Printf("%T, %T, %T, %T\n\n", i1, i2, i3, i4)

	// その他の型（まとめて宣言）
	var (
		a    float64 = 1.1
		b    string  = "hello"
		c    bool    = true
		d, e int     = 1, 2 //　2つ同時に宣言
	)
	fmt.Println(a, b, c, d, e)
	fmt.Printf("%T, %T, %T, %T, %T\n", a, b, c, d, e)

	// 少数の表示
	fmt.Printf("%f, %.2f, %e\n\n", a, a, a)

	// 型変換
	a = 1.1
	fmt.Println("そのまま表示:", a, " intに変換:", int(a))
	b = "777"
	var i int
	i, _ = strconv.Atoi(b) // 文字からintに変換 (引数2つ)
	fmt.Println("文字として表示:", b, " intに変換:", i, "\n")

	// 定数(const)の宣言
	const con = 1000
	const (
		con1 = 1000
		con2 = 2000
	)
	fmt.Println("conの値:", con)
	fmt.Println(con1, con2)

}

/* ----出力----
1 1 1 1
int, int, int, int

1.1 hello true 1 2
float64, string, bool, int, int
1.100000, 1.10, 1.100000e+00

そのまま表示: 1.1  intに変換: 1
文字として表示: 777  intに変換: 777

conの値: 1000
1000 2000
   ----------- */
