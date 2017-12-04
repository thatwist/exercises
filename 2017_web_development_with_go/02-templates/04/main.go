// Using the data structure created in the previous folder,
// modify it to hold menu information for an unlimited number of restaurants.

package main

import "html/template"
import "os"

type Restaurant struct {
	Name string
	Menu
}

type Meal struct {
	Name  string
	Price float32
}

type Menu struct {
	Breakfast []Meal
	Lunch     []Meal
	Dinner    []Meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []Restaurant{
		Restaurant{
			"Eat at Joe's",
			Menu{
				Breakfast: []Meal{Meal{"English breakfast", 3.5}, Meal{"Fruit Mix", 3.0}},
				Lunch:     []Meal{Meal{"Chicken nuggets", 5.0}, Meal{"Pizza", 4.0}},
				Dinner:    []Meal{Meal{"Pizza", 4.0}, Meal{"Cheese Mix", 2.0}}},
		},
	}

	tpl.Execute(os.Stdout, restaurants)
}
