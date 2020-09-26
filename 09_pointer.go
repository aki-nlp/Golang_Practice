package main

import (
	"fmt"
)

func fn(x int, y *int) {
	x = 2
	*y = 2
}

func main() {
	var x int = 1
	var y int = 1
	var p *int = &x

	// ポインタ
	fmt.Println(x, &x)
	fmt.Println(*p, p)

	// 値渡し, 参照渡しの違い
	fmt.Println("x:", x, "y;", y)
	fn(x, &y)
	fmt.Println("x:", x, "y;", y)

	// 領域確保(new): ポインタを返す
	var p2 *int                 // 宣言のみ(初期化しない)
	var p3 *int = new(int)      // newで初期化
	fmt.Println("p2:", " ", p2) // *pは表示できない
	fmt.Println("p3:", *p3, p3)

}

/* ----出力----
1 0xc0000b4008
1 0xc0000b4008
x: 1 y; 1
x: 1 y; 2
p2:   <nil>
p3: 0 0xc0000b4010
   ----------- */
