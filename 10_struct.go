package main

import "fmt"

// type 新しい型 元の型(typedef)
// type myInt int

type Student struct {
	id   int
	name string
}

func changeId1(s Student) {
	s.id = -1
}

func changeId2(s *Student) {
	s.id = -1
}

func main() {
	// 初期化の仕方1
	s := Student{id: 1, name: "Harry"}
	fmt.Println(s)
	fmt.Println(s.id, s.name, "\n")

	// 初期化の仕方2
	t := Student{}
	fmt.Println(t.id, t.name)
	t.id = 2
	t.name = "Ron"
	fmt.Println(t.id, t.name, "\n")

	// 初期化の仕方3
	u := Student{3, "Hermione"}
	fmt.Println(u, "\n")

	// ポインタ
	p := new(Student)
	q := &Student{}
	fmt.Println(p, q, "\n")

	// structの値渡し, 参照渡し
	fmt.Println("元の状態", u)
	changeId1(u)
	fmt.Println("changeId1を適用", u)
	changeId2(&u)
	fmt.Println("changeId2を適用", u)

}

/* ----出力----
{1 Harry}
1 Harry

0
2 Ron

{3 Hermione}

&{0 } &{0 }

元の状態 {3 Hermione}
changeId1を適用 {3 Hermione}
changeId2を適用 {-1 Hermione}
   ----------- */
