package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mr-pmillz/eoldate"
)

func main() {
	client := eoldate.NewClient()
	softwareName := "php"
	phpVersion := "7.4.33"
	isPHPVersionSupported, latestVersion, err := client.IsSupportedSoftwareVersion(softwareName, phpVersion)
	if err != nil {
		log.Fatal(err)
	}
	latestVersionInfo := fmt.Sprintf("The latest version of %s at the time of testing was %s.", strings.ToUpper(softwareName), latestVersion)
	if isPHPVersionSupported {
		fmt.Printf("%s %s is Supported. %s", softwareName, phpVersion, latestVersionInfo)
		//fmt.Sprintf("The latest version of %s at the time of testing was %s.", softwareName, latestVersion.String())
	} else {
		fmt.Printf("%s %s is no longer Supported. %s", softwareName, phpVersion, latestVersionInfo)
	}
}
