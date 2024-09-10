package main

import (
	"fmt"
	"log"

	"github.com/mr-pmillz/eoldate"
)

func main() {
	client := eoldate.NewClient()
	softwareName := "php"
	phpVersion := 8.2
	isPHPEightPointTwoSupported, err := client.IsSupportedSoftwareVersion(softwareName, phpVersion)
	if err != nil {
		log.Fatal(err)
	}
	if isPHPEightPointTwoSupported {
		fmt.Printf("%s %.1f is Supported", softwareName, phpVersion)
	} else {
		fmt.Printf("%s %.1f is not Supported", softwareName, phpVersion)
	}
}
