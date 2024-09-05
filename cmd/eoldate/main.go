package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mr-pmillz/eoldate"
	"github.com/olekukonko/tablewriter"
	"github.com/projectdiscovery/gologger"
)

func main() {
	tech := flag.String("t", "", "technology/software name to lookup")
	output := flag.String("o", "", "output directory to save results to")
	version := flag.Bool("version", false, "show version and exit")
	getAll := flag.Bool("getall", false, "get all results from all technologies")
	flag.Parse()

	eolOptions := eoldate.Options{
		Tech:    *tech,
		Output:  *output,
		Version: *version,
		GetAll:  *getAll,
	}

	if eolOptions.Version {
		fmt.Printf("Version: %s\n", eoldate.CurrentVersion)
		os.Exit(0)
	}

	if eolOptions.Output != "" {
		absOutputDir, err := eoldate.ResolveAbsPath(eolOptions.Output)
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		eolOptions.Output = absOutputDir
		if err = os.MkdirAll(absOutputDir, 0755); err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	}

	client := eoldate.NewClient()

	if eolOptions.GetAll {
		gologger.Info().Msg("Getting all available technologies")
		allProducts, err := client.GetAllProducts()
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		fmt.Println(allProducts)
		os.Exit(0)
	}

	if eolOptions.Tech == "" {
		gologger.Info().Msg("No technologies specified")
		flag.Usage()
		os.Exit(1)
	}

	releaseVersions, err := client.GetProduct(eolOptions.Tech)
	if err != nil {
		gologger.Fatal().Msgf("Error fetching product data: %v", err)
	}

	headers, includeFlags := determineHeaders(releaseVersions)
	tableString := &strings.Builder{}
	table := createTable(tableString, headers)

	currentDate := time.Now()

	for _, release := range releaseVersions {
		row := buildRow(release, includeFlags)
		colorRow(table, row, release.EOL, release.Support, currentDate)
	}

	setTableProperties(table, eolOptions.Tech, len(headers)) // Pass the number of columns
	table.Render()
	fmt.Println(tableString.String())

	if eolOptions.Output != "" {
		writeOutputFiles(eolOptions, tableString.String(), releaseVersions)
	}
}

func determineHeaders(releases []eoldate.Product) ([]string, map[string]bool) {
	headers := []string{"Cycle", "Release Date", "EOL Date", "Latest"}
	includeFlags := map[string]bool{
		"Link":                 false,
		"LatestReleaseDate":    false,
		"LTS":                  false,
		"Support":              false,
		"ExtendedSupport":      false,
		"MinJavaVersion":       false,
		"SupportedPHPVersions": false,
	}

	for _, release := range releases {
		if release.Link != "" {
			includeFlags["Link"] = true
		}
		if release.LatestReleaseDate != "" {
			includeFlags["LatestReleaseDate"] = true
		}
		if release.LTS != nil {
			includeFlags["LTS"] = true
		}
		if release.Support != nil {
			includeFlags["Support"] = true
		}
		if release.ExtendedSupport != nil {
			includeFlags["ExtendedSupport"] = true
		}
		if release.MinJavaVersion != nil {
			includeFlags["MinJavaVersion"] = true
		}
		if release.SupportedPHPVersions != "" {
			includeFlags["SupportedPHPVersions"] = true
		}
	}

	// Add headers in the desired order
	if includeFlags["Link"] {
		headers = append(headers, "Link")
	}
	if includeFlags["LatestReleaseDate"] {
		headers = append(headers, "Latest Release Date")
	}
	if includeFlags["LTS"] {
		headers = append(headers, "LTS")
	}
	if includeFlags["Support"] {
		headers = append(headers, "Support")
	}
	if includeFlags["ExtendedSupport"] {
		headers = append(headers, "Extended Support")
	}
	if includeFlags["MinJavaVersion"] {
		headers = append(headers, "Min Java Version")
	}
	if includeFlags["SupportedPHPVersions"] {
		headers = append(headers, "Supported PHP Versions")
	}

	return headers, includeFlags
}

func createTable(writer *strings.Builder, headers []string) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetHeader(headers)
	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	alignments := make([]int, len(headers))
	for i := range alignments {
		alignments[i] = tablewriter.ALIGN_CENTER
	}
	table.SetColumnAlignment(alignments)
	return table
}

func buildRow(release eoldate.Product, includeFlags map[string]bool) []string {
	row := []string{release.Cycle, release.ReleaseDate, formatEOL(release.EOL), release.Latest}

	if includeFlags["Link"] {
		row = append(row, release.Link)
	}
	if includeFlags["LatestReleaseDate"] {
		row = append(row, release.LatestReleaseDate)
	}
	if includeFlags["LTS"] {
		row = append(row, formatBool(release.LTS))
	}
	if includeFlags["Support"] {
		row = append(row, formatInterface(release.Support))
	}
	if includeFlags["ExtendedSupport"] {
		row = append(row, formatInterface(release.ExtendedSupport))
	}
	if includeFlags["MinJavaVersion"] && release.MinJavaVersion != nil {
		row = append(row, formatJavaVersion(*release.MinJavaVersion))
	}
	if includeFlags["SupportedPHPVersions"] {
		row = append(row, release.SupportedPHPVersions)
	}

	return row
}

func formatEOL(eol interface{}) string {
	switch v := eol.(type) {
	case string:
		return v
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return "N/A"
	}
}

func formatBool(value interface{}) string {
	if v, ok := value.(bool); ok {
		return fmt.Sprintf("%t", v)
	}
	return "N/A"
}

func formatInterface(value interface{}) string {
	if value == nil {
		return "N/A"
	}
	return fmt.Sprintf("%v", value)
}

func colorRow(table *tablewriter.Table, row []string, eol interface{}, support interface{}, currentDate time.Time) {
	eolDate, eolErr := parseEOLDate(eol)
	supportDate, supportErr := parseEOLDate(support)

	colors := make([]tablewriter.Colors, len(row))
	for i := range colors {
		colors[i] = tablewriter.Colors{}
	}

	if eolErr == nil && !eolDate.IsZero() {
		if eolDate.Before(currentDate) {
			colors[2] = tablewriter.Colors{tablewriter.FgRedColor}
		} else {
			colors[2] = tablewriter.Colors{tablewriter.FgGreenColor}
		}
	}

	supportIndex := -1
	for i, val := range row {
		if val == formatInterface(support) {
			supportIndex = i
			break
		}
	}

	if supportIndex != -1 && supportErr == nil && !supportDate.IsZero() {
		if supportDate.Before(currentDate) {
			colors[supportIndex] = tablewriter.Colors{tablewriter.FgRedColor}
		} else {
			colors[supportIndex] = tablewriter.Colors{tablewriter.FgGreenColor}
		}
	}

	table.Rich(row, colors)
}

func parseEOLDate(date interface{}) (time.Time, error) {
	if dateStr, ok := date.(string); ok {
		layouts := []string{
			"2006-01-02",
			"2006-01",
			"2006",
		}
		for _, layout := range layouts {
			if t, err := time.Parse(layout, dateStr); err == nil {
				return t, nil
			}
		}
	}
	return time.Time{}, fmt.Errorf("invalid date format")
}

// setTableProperties ...
func setTableProperties(table *tablewriter.Table, tech string, columnCount int) {
	footer := make([]string, columnCount)
	footer[0] = strings.ToUpper(tech)
	for i := 1; i < columnCount; i++ {
		footer[i] = ""
	}
	table.SetFooter(footer)
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
}

// Add this new function to format the Java version
func formatJavaVersion(version float64) string {
	if version == float64(int(version)) {
		return fmt.Sprintf("%.0f", version)
	}
	return fmt.Sprintf("%.1f", version)
}

func writeOutputFiles(options eoldate.Options, tableString string, releaseVersions []eoldate.Product) {
	files := map[string]func() error{
		fmt.Sprintf("%s/%s.txt", options.Output, options.Tech): func() error {
			return eoldate.WriteStringToFile(fmt.Sprintf("%s/%s.txt", options.Output, options.Tech), tableString)
		},
		fmt.Sprintf("%s/%s.json", options.Output, options.Tech): func() error {
			return eoldate.WriteStructToJSONFile(releaseVersions, fmt.Sprintf("%s/%s.json", options.Output, options.Tech))
		},
		fmt.Sprintf("%s/%s.csv", options.Output, options.Tech): func() error {
			return eoldate.WriteStructToCSVFile(releaseVersions, fmt.Sprintf("%s/%s.csv", options.Output, options.Tech))
		},
	}

	for filename, writeFunc := range files {
		if err := writeFunc(); err != nil {
			gologger.Error().Msgf("Failed to write %s: %v", filename, err)
		}
	}
}
