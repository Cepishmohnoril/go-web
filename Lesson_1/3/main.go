package main

import "fmt"

type person struct {
	fName string
}

type secretAgent struct {
	fName string
	ltk   bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fName, " person speaks")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fName, " has ltk:", sa.ltk)
}

func humanSaySomething(h human) {
	h.speak()
}

func main() {
	per := person{"Foo"}
	humanSaySomething(per)

	sa := secretAgent{"Bar", true}
	humanSaySomething(sa)
}
