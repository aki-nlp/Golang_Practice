package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func sum2(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		ch <- sum
	}
	close(ch)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go sum(s, ch)      // 15
	go sum(s[0:2], ch) // 3

	// 計算が終わった順に格納される
	x := <-ch
	y := <-ch
	fmt.Println(x, y) //3 15

	// チャネルを閉じる
	close(ch)

	// 途中計算を出力
	ch = make(chan int)
	go sum2(s, ch)
	for i := range ch {
		fmt.Println(i)
	}
}

/* ----出力----
   ----------- */
