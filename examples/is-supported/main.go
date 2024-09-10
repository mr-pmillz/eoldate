package main

import (
	"fmt"
	"log"

	"github.com/mr-pmillz/eoldate"
)

func main() {
	client := eoldate.NewClient()
	softwareName := "php"
	phpVersion := "7.4.33"
	isPHPVersionSupported, err := client.IsSupportedSoftwareVersion(softwareName, phpVersion)
	if err != nil {
		log.Fatal(err)
	}
	if isPHPVersionSupported {
		fmt.Printf("%s %s is Supported", softwareName, phpVersion)
	} else {
		fmt.Printf("%s %s is not Supported", softwareName, phpVersion)
	}
}
