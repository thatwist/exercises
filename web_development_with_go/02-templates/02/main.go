//Create a data structure to pass to a template which
//contains information about California Hotels including
//Name, Address, City, Zip, Region
//region can be: Southern, Central, Northern
//can hold an unlimited number of Hotels

package main

import "html/template"
import "os"

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	Hotels := []Hotel{
		Hotel{"Candlewood Suites Mount Pleasant", "2407 US-271 BUS", "Pleasant", 75455, "Southern"},
		Hotel{"Good Hotel", "112 7th Street", "San Francisco", 94103, "Western"},
	}
	tpl.Execute(os.Stdout, Hotels)
}
