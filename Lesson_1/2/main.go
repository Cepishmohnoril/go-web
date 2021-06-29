package main

import "fmt"

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fName, " person speaks")
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fName, " has ltk:", sa.ltk)
}

func main() {
	per := person{"Foo", "Bar"}
	sa := secretAgent{person{"Jame", "Son"}, true}
	fmt.Println(per.fName)
	per.pSpeak()
	fmt.Println(sa.ltk)
	sa.saSpeak()
	sa.person.pSpeak()
}
