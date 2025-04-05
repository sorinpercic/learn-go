package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//HTTP REQ
	res, err := http.Get("https://quotes.toscrape.com/")

	if err != nil {
		log.Fatal(err)

	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error : %d %s", res.StatusCode, res.Status)
	}
	//LOAD HTML doc
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//CREATE CSV
	file, err := os.Create("quotes.csv")
	if err != nil {
		log.Fatal("CSV file was not created, err: ", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Quote", "Author"})
	doc.Find(".quote").Each(func(index int, s *goquery.Selection) {
		quoteText := s.Find(".text").Text()
		author := s.Find(".author").Text()

		writer.Write([]string{quoteText, author})
	})
	log.Println("Saving data to CSV file")
}
