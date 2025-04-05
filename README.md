# Go Web Scraper: Quotes Scraper

This is a simple Go web scraper that collects quotes from **[Quotes to Scrape](https://quotes.toscrape.com)** and exports the data into a CSV file.

## Features
- Scrapes quotes and authors from multiple pages.
- Exports data to `quotes.csv`.
- Uses the `goquery` library to parse HTML.

## Requirements
- Go (Golang) version 1.16 or higher.
- Internet connection to access the website.

## Setup & Installation

### 1. Install Go (if not already installed)

You can download Go from [here](https://golang.org/dl/).

### 2. Install Required Go Packages

Run the following command to install `goquery`, the library used for HTML parsing:

```bash
go get github.com/PuerkitoBio/goquery
