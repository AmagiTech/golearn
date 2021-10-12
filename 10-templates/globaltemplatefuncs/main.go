package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template = nil

var fm template.FuncMap = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("sample.gohtml").Funcs(fm).ParseFiles("sample.gohtml"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "sample.gohtml", g1)

	if err != nil {
		log.Fatalln(err)
	}
}
