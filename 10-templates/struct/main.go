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
	p := person{
		Name:   "Muharrem",
		Active: true,
	}

	err := tpl.Execute(os.Stdout, p)

	if err != nil {
		log.Fatalln(err)
	}
}
