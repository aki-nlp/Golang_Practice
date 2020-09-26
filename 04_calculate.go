package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1
	y := 3

	// 数値計算
	fmt.Printf("x = %d  y = %d\n", x, y)
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x / y = %f\n", float64(x)/float64(y))
	fmt.Printf("y %% 2 = %d\n", y%2)
	fmt.Printf("x < y %v\n\n", x < y)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("true || false = %v\n\n", true || false)

	// インクリメント, デクリメント
	i := 0
	fmt.Println(i)
	i++
	fmt.Println(i)
	i += 2
	fmt.Println(i)
	i--
	fmt.Println(i, "\n")

	// math.Powは, 引数, 返り値ともにfloat64
	fmt.Printf("yの3乗: %f\n", math.Pow(float64(y), 3))
	fmt.Printf("ルートy: %f\n\n", math.Sqrt(float64(y)))

	// 切り上げ, 切り捨て, 四捨五入
	z := 1.732051
	fmt.Printf("切り上げ: %f\n", math.Ceil(z))
	fmt.Printf("切り捨て: %f\n", math.Floor(z))
	fmt.Printf("四捨五入: %f\n", math.Round(z))
}

/* ----出力----
x = 1  y = 3
x + y = 4
x - y = -2
x * y = 3
x / y = 0
x / y = 0.333333
y % 2 = 1
x < y true

true && false = false
true || false = true

0
1
3
2

yの3乗: 27.000000
ルートy: 1.732051

切り上げ: 2.000000
切り捨て: 1.000000
四捨五入: 2.000000
   ----------- */
