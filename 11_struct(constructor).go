package main

import "fmt"

type Student struct {
	id      int
	classID int
	name    string
}

// 値を書き換えないメソッド
func (s Student) getClassID() int {
	return ((s.id - 1) % 3) + 1
}

// 値を書き換えるメソッド
func (s *Student) changeClassID() {
	s.classID = ((s.id + 1) % 3) + 1
}

// コンストラクタ
func NewStudent(id int, classID int, name string) *Student {
	return &Student{id, classID, name}

	// 下の書き方でもいい
	// s := new(Student)
	// s.id = id
	// s.classID = classID
	// s.name = name
	// return s
}

// Stringers(Println(s)の出力を変える, Pythonの__str__のようなもの)
func (s Student) String() string {
	return "My name is " + s.name
}

func main() {
	s := Student{id: 1, name: "Harry"}
	fmt.Println("classID:", s.getClassID())
	s.changeClassID()
	fmt.Println(s.name+"'s", "class has changed to", s.classID)

	t := NewStudent(2, 0, "Ron")
	fmt.Println(t)
	t.changeClassID()
	fmt.Println(t.name+"'s", "class has changed to", t.classID)

}

/* ----出力----
classID: 1
Harry's class has changed to 3
My name is Ron
Ron's class has changed to 1
   ----------- */
