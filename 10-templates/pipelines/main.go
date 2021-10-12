package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template = nil

func double(x int) int {
	return x * 2
}

func square(x int) float64 {
	return float64(x * x)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"db":   double,
	"sq":   square,
	"sqrt": squareRoot,
}

func init() {
	tpl = template.Must(template.New("sample.gohtml").Funcs(fm).ParseFiles("sample.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "sample.gohtml", 3)

	if err != nil {
		log.Fatalln(err)
	}
}
