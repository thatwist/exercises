// Create a data structure to pass to a template which
// contains information about restaurant's menu
// including Breakfast, Lunch, and Dinner items

package main

import "html/template"
import "os"

type Meal struct {
	Name  string
	Price float32
}

type menu struct {
	Breakfast []Meal
	Lunch     []Meal
	Dinner    []Meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		Breakfast: []Meal{Meal{"English breakfast", 3.5}, Meal{"Fruit Mix", 3.0}},
		Lunch:     []Meal{Meal{"Chicken nuggets", 5.0}, Meal{"Pizza", 4.0}},
		Dinner:    []Meal{Meal{"Pizza", 4.0}, Meal{"Cheese Mix", 2.0}},
	}

	tpl.Execute(os.Stdout, m)
}
