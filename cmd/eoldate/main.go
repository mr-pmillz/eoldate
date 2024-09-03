package main

import (
	"flag"
	"fmt"
	"github.com/mr-pmillz/eoldate"
	"github.com/olekukonko/tablewriter"
	"github.com/projectdiscovery/gologger"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	tech := flag.String("t", "", "technology/software name to lookup")
	output := flag.String("o", "", "output directory to save results to")
	version := flag.Bool("version", false, "show version and exit")
	getAll := flag.Bool("getall", false, "get all results from all technologies")
	flag.Parse()

	eolOptions := eoldate.Options{}
	if *output != "" {
		absOutputDir, err := eoldate.ResolveAbsPath(*output)
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		eolOptions.Output = absOutputDir
		if err = os.MkdirAll(absOutputDir, 0755); err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	}

	eolOptions.Tech = *tech
	eolOptions.Version = *version
	eolOptions.GetAll = *getAll

	if eolOptions.Version {
		fmt.Printf("Version: %s\n", eoldate.CurrentVersion)
		os.Exit(0)
	}

	client := eoldate.NewClient(eoldate.EOLBaseURL)

	if *getAll {
		gologger.Info().Msg("Getting all available technologies")
		allProducts, err := client.GetAllProducts()
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		fmt.Println(allProducts)
		os.Exit(0)
	}

	releaseVersions, err := client.GetProduct(eolOptions.Tech)
	if err != nil {
		fmt.Printf("Error fetching product data: %v\n", err)
		return
	}
	// Create a new table
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Cycle", "Release Date", "EOL Date", "Latest", "Latest Release Date", "LTS", "Support"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor},
	)
	// Get current date
	currentDate := time.Now()

	row := make([]string, 0)
	var eolDateTime time.Time
	// Add data to the table
	for _, release := range releaseVersions {
		var EOLIsBool bool
		rt := reflect.TypeOf(release.EOL)
		if rt.Kind() == reflect.String {
			eolDateTime, err = time.Parse("2006-01-02", release.EOL.(string))
			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue
			}
			row = []string{release.Cycle, release.ReleaseDate, release.EOL.(string), release.Latest, release.LatestReleaseDate, fmt.Sprintf("%t", release.LTS), release.Support}
		}
		if rt.Kind() == reflect.Bool {
			EOLIsBool = true
			row = []string{release.Cycle, release.ReleaseDate, "N/A", release.Latest, release.LatestReleaseDate, fmt.Sprintf("%t", release.LTS), release.Support}
		}

		// Check if EOL date is older or later than the current date
		if eolDateTime.Before(currentDate) && !EOLIsBool {
			table.Rich(row, []tablewriter.Colors{{}, {}, tablewriter.Colors{tablewriter.FgRedColor}, {}, {}, {}, {}})
		} else {
			table.Rich(row, []tablewriter.Colors{{}, {}, tablewriter.Colors{tablewriter.FgGreenColor}, {}, {}, {}, {}})
		}
	}

	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
	})

	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	table.SetFooter([]string{
		strings.ToUpper(eolOptions.Tech),
		"",
		"",
		"",
		"",
		"",
		"",
	})
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)

	// Render the table
	table.Render()
	fmt.Println(tableString.String())

	if eolOptions.Output != "" {
		tableOutputFile := fmt.Sprintf("%s/%s.txt", eolOptions.Output, eolOptions.Tech)
		softwareEolDateCSV := fmt.Sprintf("%s/%s.csv", eolOptions.Output, eolOptions.Tech)
		softwareEolDateJSON := fmt.Sprintf("%s/%s.json", eolOptions.Output, eolOptions.Tech)

		if err = eoldate.WriteStringToFile(tableOutputFile, tableString.String()); err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		if err = eoldate.WriteStructToJSONFile(releaseVersions, softwareEolDateJSON); err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		if err = eoldate.WriteStructToCSVFile(releaseVersions, softwareEolDateCSV); err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	}
}
