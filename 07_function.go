package main

import "fmt"

// 返り値は引数の後に指定
func add(x int, y int) int {
	return x + y
}

// 返り値の変数を最初にに指定
func minus(x int, y int) (result int) {
	result = x - y
	return // 最初に宣言しているので「return result」と書かなくていい
}

// 引数が同じ型ならまとめて書ける
// 返り値が2つなら引数の後に2つ指定する
func addMinus(x, y int) (int, int) {
	return x + y, x - y
}

// 可変引数
func funcA(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}

// クロージャ(関数閉方): 外側の変数を記憶した関数
// この関数の場合, piを記憶
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(minus(1, 2))
	a, b := addMinus(1, 2)
	fmt.Println(a, b, "\n")

	// 関数内で関数を宣言
	f := func(x int) {
		fmt.Println("引数は", x)
	}
	f(1000)

	// 関数に名前をつけないで実行
	func(x int) {
		fmt.Println("引数は", x)
	}(10)
	fmt.Println()

	// 可変引数
	funcA(1, 2, 3, 4)
	x := []int{1, 2, 3, 4}
	funcA(0, x...)
	fmt.Println()

	// クロージャ
	c1 := circleArea(3.14)
	fmt.Println("半径1の円の面積:", c1(1))
	fmt.Println("半径2の円の面積:", c1(2))

	c2 := circleArea(3)
	fmt.Println("半径1の円の面積:", c2(1))
	fmt.Println("半径2の円の面積:", c2(2))
}

/* ----出力----
3
-1
3 -1

引数は 1000
引数は 10

a=1
b[0]=2
b[1]=3
b[2]=4
a=0
b[0]=1
b[1]=2
b[2]=3
b[3]=4

半径1の円の面積: 3.14
半径2の円の面積: 12.56
半径1の円の面積: 3
半径2の円の面積: 12
   ----------- */
