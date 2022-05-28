package main

import (
	"fmt"
	"github.com/muesli/regommend"
)

func main() {
	// Accessing a new regommend table for the first time will create it.
	books := regommend.Table("books")

	booksChrisRead := make(map[interface{}]float64)
	booksChrisRead["1984"] = 5.0
	booksChrisRead["Robinson Crusoe"] = 5.0
	booksChrisRead["Moby-Dick"] = 3.0
	books.Add("Chris", booksChrisRead)

	booksJaRead := make(map[interface{}]float64)
	booksJaRead["1984"] = 5.0
	booksJaRead["Robinson Crusoe"] = 5.0
	booksJaRead["GOG"] = 7.0
	booksJaRead["Gulliver's Travels"] = 3.5
	books.Add("Ja", booksJaRead)

	booksJayRead := make(map[interface{}]float64)
	booksJayRead["1984"] = 1.0
	booksJayRead["Robinson Crusoe"] = 1.0
	booksJayRead["Robinson"] = 4.0
	booksJayRead["Gulliver's Travels"] = 1.5
	books.Add("Jay", booksJayRead)

	recs, _ := books.Recommend("Chris")
	for _, rec := range recs {
		fmt.Println("Recommending", rec.Key, "with score", rec.Distance)
	}

	neighbors, _ := books.Neighbors("Chris")
	for _, n := range neighbors {
		fmt.Println("Recommending", n.Key, "with score", n.Distance)
	}
}
