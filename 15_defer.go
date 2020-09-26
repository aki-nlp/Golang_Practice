package main

import (
	"fmt"
	"os"
)

func f1() {
	fp, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("The file can't open!")
	}
	defer fp.Close()

	defer fmt.Println("The third message!")
	fmt.Println("The second message!")
}

func main() {
	// deferは遅延処理
	defer fmt.Println("The final message!")
	defer fmt.Println("The fifth message!")
	defer fmt.Println("The forth message!")
	fmt.Println("The first message!")
	f1()
}

/* ----出力----
The first message!
The file can't open!
The second message!
The third message!
The forth message!
The fifth message!
The final message!
   ----------- */
