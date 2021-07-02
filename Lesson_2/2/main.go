package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	name, address, city, zip, region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotelsList := []hotel{
		{
			name:    "Doot",
			address: "Some Address",
			city:    "New Foo",
			zip:     "Bar",
			region:  "Memeland",
		},
	}

	err := tpl.Execute(os.Stdout, hotelsList)

	if err != nil {
		log.Fatalln(err)
	}
}
