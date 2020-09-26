package main

import "fmt"

type Human interface {
	useMagic()
}

type Wizardkind struct {
	name string
}

type Muggle struct {
	name string
}

func (w Wizardkind) useMagic() {
	fmt.Println(w.name, "used", "\"Expecto patronum!!!\"")
}

func (m Muggle) useMagic() {
	fmt.Println(m.name, "can't use any magic!")
}

func main() {
	var harry Human = Wizardkind{"Harry"}
	var dudley Human = Muggle{"Dudley"}
	harry.useMagic()
	dudley.useMagic()
}

/* ----出力----
Harry used "Expecto patronum!!!"
Dudley can't use any magic!
   ----------- */
