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

func main() {
	colors := map[string]string{
		"Red":   "Kırmızı",
		"Blue":  "Mavi",
		"Green": "Yeşil",
	}

	err := tpl.Execute(os.Stdout, colors)

	if err != nil {
		log.Fatalln(err)
	}
}
