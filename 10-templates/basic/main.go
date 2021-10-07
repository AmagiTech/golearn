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
	err = tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", "Muharrem")

	if err != nil {
		log.Fatalln(err)
	}

}
