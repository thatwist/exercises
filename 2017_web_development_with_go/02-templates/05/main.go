// 1. Parse this CSV file, putting two fields from the contents of
//    the CSV file into a data structure.
// 2. Parse an html template, then pass the data from step 1 into
//    the CSV template; have the html template display the CSV data in a web page.

package main

import (
	"bufio"
	"encoding/csv"
	"html/template"
	"os"
)

type Recording struct {
	Date string
	Open string
}

var (
	recordings []Recording
	tpl        *template.Template
)

func init() {
	file, err := os.Open("table.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	recordings = make([]Recording, len(records))
	for i, r := range records {
		recordings[i] = Recording{r[0], r[1]}
	}
}

func main() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl.Execute(os.Stdout, recordings)
}
