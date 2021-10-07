package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")

	if err != nil {
		log.Fatalln(err)
	}

	arr := []string{"A", "B", "C", "D"}
	err = tpl.ExecuteTemplate(os.Stdout, "tmp.gohtml", arr)

	if err != nil {
		log.Fatalln(err)
	}

}
