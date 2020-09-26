package main

import "fmt"

type Student struct {
	id   int
	name string
}

type SeniorStudent struct {
	Student
	teacher string
}

func NewSeniorStudent(id int, name string, teacher string) *SeniorStudent {
	return &SeniorStudent{Student{id, name}, teacher}
}

func (s Student) printName() {
	fmt.Println(s.name)
}

func (g SeniorStudent) printTeacher() {
	fmt.Println(g.teacher)
}

func (g *SeniorStudent) changeTeacher(teacher string) {
	g.teacher = teacher
}

func main() {
	g := NewSeniorStudent(1, "Harry", "Snape")
	fmt.Println(g)
	g.printName()
	g.printTeacher()
	g.changeTeacher("McGonagall")
	g.printTeacher()

}

/* ----出力----
&{{1 Harry} Snape}
Harry
Snape
McGonagall
   ----------- */
