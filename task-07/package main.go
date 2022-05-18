package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "billionaresdata.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "age", "NetWorth", "Source", "Countryterr"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".name"),
			e.ChildText(".age"),
			e.ChildText(".Net-Worth"),
			e.ChildText(".source"),
			e.ChildText(".Country/Territory"),
		})
	})

	c.Visit("https://www.forbes.com/real-time-billionaires/#355c2f3f3d78")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
