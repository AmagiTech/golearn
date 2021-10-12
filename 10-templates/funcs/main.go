package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func firstThreeLetter(p string) string {
	return strings.TrimSpace(p)[:3]
}

var fmap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThreeLetter,
}

var tpl *template.Template = nil

type car struct {
	Brand string
	Model int
}

func init() {
	tpl = template.Must(template.New("sample.gohtml").Funcs(fmap).ParseFiles("sample.gohtml"))
}

func main() {
	a := car{
		Brand: "Ford",
		Model: 1992,
	}

	b := car{
		Brand: "Toyota",
		Model: 2000,
	}

	vm := []car{a, b}

	err := tpl.Execute(os.Stdout, vm)

	if err != nil {
		log.Fatalln(err)
	}
}
