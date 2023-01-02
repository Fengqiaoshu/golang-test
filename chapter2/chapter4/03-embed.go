package main

import "fmt"

// 定义结构体
type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

// 注意：是否可有使用非*号呢
func (u *Person) setAge(_age int) {
	u.age = _age

}

type SuperMan struct {
	Person   //继承
	strength int
	speed    int
}

func (s *SuperMan) print() {
	fmt.Printf(" name = %s,age = %d, sex = %s, fight = %d\n",
		s.name, s.age, s.sex, s.fight)
	fmt.Printf("stength = %d, speed = %d\n", s.strength, s.speed)
}

func main() {
	s1 := SuperMan{
		Person: Person{
			name:  "clark",
			age:   29,
			sex:   "man",
			fight: 200000,
		},
		strength: 1000,
		speed:    2313541,
	}

	s1.print()
	s1.setAge(35)
	s1.print()
}
