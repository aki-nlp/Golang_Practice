package main

import (
	"fmt"
	"sync"
	"time"
)

func funcA() {
	for i := 0; i < 5; i++ {
		fmt.Println("funcA")
		time.Sleep(1000 * time.Millisecond)
	}
}

func funcA2(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("funcA")
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}

func funcB() {
	for i := 0; i < 5; i++ {
		fmt.Println("funcB")
		time.Sleep(1000 * time.Millisecond)
	}
}

func funcC() {
	for i := 0; i < 5; i++ {
		fmt.Println("funcC")
		// time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// Aを並列化
	go funcA()
	funcB()

	fmt.Println("")

	// Aを並列化するが先にCが終了してしまいAが5個出力されない
	go funcA()
	funcC()

	fmt.Println("")
	// 並列化したfuncAの出力を待つ
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("")

	// 並列処理が終わるのを待つ
	var wg sync.WaitGroup
	wg.Add(1)
	go funcA2(&wg)
	funcC()
	wg.Wait()

}

/* ----出力----
funcB
funcA
funcA
funcB
funcB
funcA
funcA
funcB
funcB
funcA

funcC
funcA
funcC
funcC
funcC
funcC

funcA
funcA
funcA
funcA

funcC
funcC
funcC
funcC
funcC
funcA
funcA
funcA
funcA
funcA
   ----------- */
