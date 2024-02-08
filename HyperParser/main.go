package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const URL = "https://hypeauditor.com/top-instagram-all-russia/"

type Data struct {
	Rank        string
	Subscribers string
	Name        string
	Title       string
	Category    string
}

func main() {
	var result []Data
	c := colly.NewCollector()

	c.OnHTML(".table", func(e *colly.HTMLElement) {
		e.ForEach(".row", func(_ int, e *colly.HTMLElement) {
			row := Data{
				Rank:        e.ChildText(".rank > span"),
				Subscribers: e.ChildText(".subscribers"),
				Name:        e.ChildText(".contributor__name-content"),
				Title:       e.ChildText(".contributor__title"),
				Category:    e.ChildText(".category"),
			}
			result = append(result, row)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		saveToCSV(result, "output.csv")
		log.Println("Scraping completed")
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error occurred:", err)
	})

	c.Visit(URL)
}

func saveToCSV(data []Data, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()

	// For UTF8
	file.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Rank", "Subscribers", "Name", "Title", "Category"})

	for _, value := range data {
		str := []string{value.Rank, value.Subscribers, value.Name, value.Title, value.Category}
		err := writer.Write(str)
		if err != nil {
			log.Println("Failed to write to CSV:", err)
		}
	}

	log.SetOutput(os.Stdout)
	log.Println("Data saved to", filename)
}
