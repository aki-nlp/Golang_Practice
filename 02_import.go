package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	// 時刻の表示
	fmt.Println(time.Now())

	// ユーザ情報を表示
	fmt.Println(user.Current())
}

/* ----出力----
3
-1
3 -1

引数は 1000
引数は 10
   ----------- */
