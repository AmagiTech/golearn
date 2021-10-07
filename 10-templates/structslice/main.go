package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("sample.gohtml"))
}

type person struct {
	Name   string
	Active bool
}

func main() {
	p1 := person{
		Name:   "Muharrem",
		Active: true,
	}

	p2 := person{
		Name:   "Ali",
		Active: false,
	}

	err := tpl.Execute(os.Stdout, []person{p1, p2})

	if err != nil {
		log.Fatalln(err)
	}
}
