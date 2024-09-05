package main

import (
	"fmt"
	"github.com/mr-pmillz/eoldate"
	"log"
)

func main() {
	client := eoldate.NewClient()
	products, err := client.GetProduct("php")
	if err != nil {
		log.Fatalf("Error fetching product data: %v", err)
	}

	versionToCheck := 7.4

	for _, product := range products {
		supported, err := product.IsVersionSupported(versionToCheck)
		if err != nil {
			continue
		}

		if supported {
			fmt.Printf("PHP %.1f is still supported\n", versionToCheck)
		} else {
			fmt.Printf("PHP %.1f is no longer supported\n", versionToCheck)
		}
		return // Exit after finding the matching cycle
	}

	fmt.Printf("PHP %.1f was not found in any product cycle\n", versionToCheck)
}
