package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(i, "is", "int")
	case string:
		fmt.Println(i, "is", "string")
	default:
		fmt.Printf("%v is %T\n", i, v)
	}
}

func main() {
	do(100)
	do("hello")
	do(true)
	fmt.Println("")

	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

/* ----出力----
3
-1
3 -1

引数は 1000
引数は 10
   ----------- */
