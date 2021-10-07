package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("sample.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	arr := []string{"A", "B", "C", "D"}
	err = tpl.Execute(os.Stdout, arr)

	if err != nil {
		log.Fatalln(err)
	}

}
