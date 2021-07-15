package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	//http.HandleFunc("/", index)
	//http.HandleFunc("/dog", dog)
	//http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln("Error parsing tpl", err)
	}

	err = tpl.Execute(w, "Index")
	if err != nil {
		log.Fatalln("Error executing tpl", err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln("Error parsing tpl", err)
	}

	err = tpl.Execute(w, "Dog")
	if err != nil {
		log.Fatalln("Error executing tpl", err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln("Error parsing tpl", err)
	}

	err = tpl.Execute(w, "Me")
	if err != nil {
		log.Fatalln("Error executing tpl", err)
	}
}
