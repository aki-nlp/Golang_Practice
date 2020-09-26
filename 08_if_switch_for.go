package main

import "fmt"

// if文の書き方
func funcIf(n int) {
	fmt.Printf("n=%2d -> ", n)
	if n < 0 {
		fmt.Println("n < 0")
	} else if n == 0 {
		fmt.Println("n = 0")
	} else {
		fmt.Println("n > 0")
	}
}

// switch文の書き方1
func funcSwitch1(n int) {
	fmt.Printf("n=%2d -> ", n)
	switch {
	case n < 0:
		fmt.Println("n < 0")
	case n == 0:
		fmt.Println("n = 0")
	default:
		fmt.Println("n > 0")
	}
}

// switch文の書き方2
func funcSwitch2() {
	fmt.Printf("The size is ")
	size := "L"
	switch size {
	case "S":
		fmt.Println("S.")
	case "L":
		fmt.Println("L.")
	default:
		fmt.Println("M.")
	}
}

// for文の書き方(golangにwhileはない)
func funcFor1(n int) {
	for i := 0; i < n; i++ {
		if 3 <= i && i <= 5 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println(i)
	}
}

// while風な書き方
func funcFor2(n int) {
	for 0 < n {
		fmt.Println("hello")
		n--
	}
}

// rangeを使った書き方
func funcFor3() {
	// スライス
	cities := []string{"Tokyo", "Osaka", "Nagoya"}
	for i, city := range cities {
		fmt.Println(i, city)
	}

	// map
	m := map[string]int{
		"x": 1,
		"y": 2,
	}
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

}

func main() {
	fmt.Println("[FuncIf]")
	funcIf(-1)
	funcIf(0)
	funcIf(2)

	fmt.Println("\n" + "[funcSwitch1]")
	funcSwitch1(-1)
	funcSwitch1(0)
	funcSwitch1(2)

	fmt.Println("\n" + "[funcSwitch2]")
	funcSwitch2()

	fmt.Println("\n" + "[FuncFor1]")
	funcFor1(10)

	fmt.Println("\n" + "[FuncFor2]")
	funcFor2(2)

	fmt.Println("\n" + "[FuncFor3]")
	funcFor3()
}

/* ----出力----
[FuncIf]
n=-1 -> n < 0
n= 0 -> n = 0
n= 2 -> n > 0

[funcSwitch1]
n=-1 -> n < 0
n= 0 -> n = 0
n= 2 -> n > 0

[funcSwitch2]
The size is L.

[FuncFor1]
0
1
2
6
7

[FuncFor2]
hello
hello

[FuncFor3]
0 Tokyo
1 Osaka
2 Nagoya
y : 2
x : 1
   ----------- */
