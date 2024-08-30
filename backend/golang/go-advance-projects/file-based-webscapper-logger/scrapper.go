package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func Scrapper() {
	c := colly.NewCollector()

	// h is the current html element
	c.OnHTML(".product", func(h *colly.HTMLElement) {

		p := Product{}

		p.Url = h.ChildAttr("a", "href")
		p.Image = h.ChildAttr("img", "src")
		p.Title = h.ChildText(".product-name")
		p.Price = h.ChildText(".price")

		products = append(products, p)
	})

	c.OnScraped(func(r *colly.Response) {
		CSVgen()
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}

func CSVgen() {

	// Create a csv file
	file, err := os.Create("product.csv")
	if err != nil {
		log.Printf("failed to create file :- %v", err.Error())
		return
	}
	defer file.Close()

	// CSV writer
	w := csv.NewWriter(file)
	defer w.Flush()

	// csv header column
	header := []string{
		"Url",
		"Image",
		"Title",
		"Price",
	}

	w.Write(header)

	for i, product := range products {

		// Rows for csv files
		record := []string{
			product.Url,
			product.Image,
			product.Title,
			product.Price,
		}

		w.Write(record)

		log.Printf("record entered %v", i)
	}

}

// li -> .product
// childAttr a href
// childAttr img src
// childText .product-name
// childText .price
