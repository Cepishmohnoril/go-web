package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotelsList := []hotel{
		{
			Name:    "Doot",
			Address: "Some Address",
			City:    "New Foo",
			Zip:     "Bar",
			Region:  "Memeland",
		},
		{
			Name:    "Doot2",
			Address: "Some Address2",
			City:    "New Bao",
			Zip:     "For",
			Region:  "Memeland",
		},
	}

	err := tpl.Execute(os.Stdout, hotelsList)

	if err != nil {
		log.Fatalln(err)
	}
}
