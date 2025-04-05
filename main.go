package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.Create("quotes.csv")
	if err != nil {
		log.Fatal("CSV file was not created, err: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Quote", "Author"})

	page := 1

	for {
		url := "https://quotes.toscrape.com/page/" + strconv.Itoa(page) + "/"

		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatal(err)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".quote").Each(func(i int, s *goquery.Selection) {

			quoteText := s.Find(".text").Text()
			author := s.Find(".author").Text()

			writer.Write([]string{quoteText, author})
		})
		nextPage := doc.Find(".next a").AttrOr("href", "")
		if nextPage == "" {
			break
		}
		page++
	}
	log.Println("Everything works fine... ")
}
