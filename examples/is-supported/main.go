package main

import (
	"fmt"
	"github.com/mr-pmillz/eoldate"
	"log"
	"time"
)

func main() {
	client := eoldate.NewClient()
	softwareName := "php"
	phpVersion := "7.4"
	isSupported, latestVersion, product, err := client.IsSupportedSoftwareVersion(softwareName, phpVersion)
	if err != nil {
		log.Fatal(err)
	}

	latestVersionInfo := fmt.Sprintf("The latest version of %s on %v was %s.", softwareName, time.Now().Format("01-02-2006"), latestVersion)

	endDate := product.GetEndDate()
	if isSupported {
		fmt.Printf("%s %s is Supported. %s\n", softwareName, phpVersion, latestVersionInfo)

		if endDate != nil {
			years, months, days := eoldate.CalculateTimeDifference(*endDate)
			fmt.Printf("Support ends in %d years, %d months, and %d days (%s)\n", years, months, days, endDate.Format("01-02-2006"))
		}
	} else {
		fmt.Printf("%s %s is no longer Supported. %s\n", softwareName, phpVersion, latestVersionInfo)

		if endDate != nil {
			years, months, days := eoldate.CalculateTimeDifference(*endDate)
			fmt.Printf("Support ended %d years, %d months, and %d days ago (%s)\n", years, months, days, endDate.Format("01-02-2006"))
		}
	}
}
