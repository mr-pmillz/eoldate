package main

import (
	"fmt"
	"log"

	"github.com/mr-pmillz/eoldate"
)

func main() {
	client := eoldate.NewClient()
	products, err := client.GetProduct("php")
	if err != nil {
		log.Fatalf("Error fetching product data: %v", err)
	}

	versionsToCheck := []float64{5.6, 7.4, 8.0, 8.1, 8.2}

	for _, version := range versionsToCheck {
		supported, err := products.IsVersionSupported(version)
		if err != nil {
			fmt.Printf("PHP %.1f: %v\n", version, err)
			continue
		}

		if supported {
			fmt.Printf("PHP %.1f is still supported\n", version)
		} else {
			fmt.Printf("PHP %.1f is no longer supported\n", version)
		}
	}
}
