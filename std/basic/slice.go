package basic

import "log"

func Slice1() {
	s1 := []string{"a,b,c,d"}
	for _, v := range s1 {
		v = v + v
	}
	//[a,b,c,d]
	log.Println(s1)

	s2 := &[]string{"a,b,c,d"}
	for _, v := range *s2 {
		v = v + v
	}
	//[a,b,c,d]
	log.Println(s1)

}

func Slice2() {
	c1 := []*Coder{&coder}
	for _, v := range c1 {
		v.Name = "hello kitty"
	}
	log.Println(c1[0].Name)

	c2 := []interface{}{&coder}
	for _, v := range c2 {
		v.(*Coder).Name = "hello zing"
	}
	log.Println(c2[0].(*Coder).Name)
}

type F func(str string)
type Slice struct {
	name string
	f    F
}

func (s Slice) run() {
	s.f(s.name)
}

func Slice3() {
	s1 := []Slice{{name: "zing"}, {name: "golang"}}
	for _, v := range s1 {
		v.f = func(str string) {
			log.Println(str)
		}
	}
	//invalid memory address or nil pointer dereference [recovered]
	//s1[0].run()

	s2 := []*Slice{{name: "zing"}, {name: "golang"}}
	for _, v := range s2 {
		v.f = func(str string) {
			log.Println(str)
		}
	}
	s2[0].run()
}

func arr(i int) {
	a := []int{1, 2, 3, 4}
	for i, v := range a {
		if v == i {
			//delete(a, i)
		}
	}

}
func Slice4() {
	/*	arr := []int{1, 2, 3, 4, 5}

		for _, v := range arr {

		}*/
}
