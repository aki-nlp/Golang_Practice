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
