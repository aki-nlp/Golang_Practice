package main

import (
	"fmt"
)

func main() {

	// 配列の宣言(サイズ変更(appendなど)できない)
	var a [5]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	fmt.Println("a:", a)

	var b [5]int = [5]int{0, 1, 2}
	fmt.Println("b:", b, "\n")

	// スライスの宣言
	var c []int = []int{0, 1, 2}
	fmt.Println("c:", c)
	c = append(c, 3) // 要素の追加
	fmt.Println("c:", c)
	fmt.Println("c[2]:", c[2])
	fmt.Println("c[2:]:", c[2:])
	fmt.Println("len(c):", len(c), "\n")

	// 2次元配列(スライス)
	var d = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println("d:", d)
	fmt.Println("d[2][2]:", d[2][2])

	// makeとcapacity (メモリの話が関わってくる)
	e := make([]int, 5)
	f := make([]int, 0, 5) // make([]int, length, capacity)
	for i := 0; i < 5; i++ {
		e = append(e, i)
		f = append(f, i)
	}
	fmt.Println("e:", e)
	fmt.Println("f:", f, "\n")

	// map(Pythonの辞書型のようなもの)
	m := map[string]int{
		"x": 1,
		"y": 2,
	}
	m["z"] = 3     // 要素の追加
	m["a"] = 4     // 要素の追加
	delete(m, "a") // 要素の削除
	fmt.Println("m:", m)
	fmt.Println("m[\"x\"]:", m["x"])
	v, ok := m["y"] // 要素が含まれるか
	fmt.Println("v:", v, "ok:", ok)
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
	fmt.Println("")

	// 空のmapの宣言
	n := map[string]int{}
	// n := make(map[string]int)
	n["a"] = 1
	fmt.Println("n:", n, "\n")

}

/* ----出力----
a: [0 1 2 0 0]
b: [0 1 2 0 0]

c: [0 1 2]
c: [0 1 2 3]
c[2]: 2
c[2:]: [2 3]
len(c): 4

d: [[1 2 3] [4 5 6] [7 8 9]]
d[2][2]: 9
e: [0 0 0 0 0 0 1 2 3 4]
f: [0 1 2 3 4]

m: map[x:1 y:2 z:3]
m["x"]: 1
v: 2 ok: true
x : 1
y : 2
z : 3

n: map[a:1]
   ----------- */
