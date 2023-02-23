package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	ID    string
	Name  string
	Price float64
}

func main() {

	// Open the CSV file
	file, err := os.Open("items.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	items := parseCSV(file)

	// Output the data
	fmt.Printf("%-10s %-20s %10s\n", "ID", "Name", "Price")
	for _, item := range items {
		fmt.Printf("%-10s %-20s $%7.2f\n", item.ID, item.Name, item.Price)
	}
}

// Parse the CSV file
func parseCSV(file *os.File) []Item {
	reader := csv.NewReader(file)
	var items []Item
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		price := 0.0
		_, err = fmt.Sscanf(record[2], "%f", &price)
		if err != nil {
			log.Fatal(err)
		}

		item := Item{
			ID:    record[0],
			Name:  record[1],
			Price: price,
		}
		items = append(items, item)
	}

	return items

}
