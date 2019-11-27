package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// csv 简单实用
func main() {
	//WriterCsvExample()
	ReaderCsvExample()
}

// csv写入
func WriterCsvExample() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	file, _ := os.Create("chapter07/e.csv")
	w := csv.NewWriter(file)
	defer file.Close()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// csv 读取
func ReaderCsvExample() {
	//	in := `first_name,last_name,username
	//"Rob","Pike",rob
	//Ken,Thompson,ken
	//"Robert","Griesemer","gri"`
	file, _ := os.Open("chapter07/e.csv")
	//reader := csv.NewReader(strings.NewReader(in))
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(record)
}

// fuzz csv
func FuzzCsvExample() {

}
