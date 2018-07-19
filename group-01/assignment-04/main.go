package main

import (
    "log"
	"os"
    "encoding/csv"
    "text/template"
	"io"
	"bufio"
)

type tradeDay struct {
	Date string
	Open string
	Close string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	csvFile, _ := os.Open("table.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var tradeDays []tradeDay
	
	for {    
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        
        tradeDays = append(tradeDays, tradeDay{
            Date: line[0],
            Open: line[1],
            Close: line[4],
        })
    }

	err := tpl.Execute(os.Stdout, tradeDays)
	if err != nil {
		log.Fatalln(err)
	}
}