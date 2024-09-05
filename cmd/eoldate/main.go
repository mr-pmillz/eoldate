package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/mr-pmillz/eoldate"
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

	data, err := client.Get(fmt.Sprintf("%s.json", eolOptions.Tech))
	if err != nil {
		gologger.Fatal().Msgf("Error fetching product data: %v", err)
	}

	var products []eoldate.Product
	if err = json.Unmarshal(data, &products); err != nil {
		gologger.Fatal().Msgf("Error parsing JSON: %v", err)
	}

	tableBuilder := NewTableBuilder(products)
	tableString := tableBuilder.Render()
	fmt.Println(tableString)

	if eolOptions.Output != "" {
		writeOutputFiles(eolOptions, tableString, products)
	}
}

func writeOutputFiles(options eoldate.Options, tableString string, products []eoldate.Product) {
	files := map[string]func() error{
		fmt.Sprintf("%s/%s.txt", options.Output, options.Tech): func() error {
			return eoldate.WriteStringToFile(fmt.Sprintf("%s/%s.txt", options.Output, options.Tech), tableString)
		},
		fmt.Sprintf("%s/%s.json", options.Output, options.Tech): func() error {
			return eoldate.WriteStructToJSONFile(products, fmt.Sprintf("%s/%s.json", options.Output, options.Tech))
		},
		fmt.Sprintf("%s/%s.csv", options.Output, options.Tech): func() error {
			return eoldate.WriteStructToCSVFile(products, fmt.Sprintf("%s/%s.csv", options.Output, options.Tech))
		},
	}

	for filename, writeFunc := range files {
		if err := writeFunc(); err != nil {
			gologger.Error().Msgf("Failed to write %s: %v", filename, err)
		}
	}
}

// TableBuilder handles the creation and population of the table
type TableBuilder struct {
	products []eoldate.Product
	headers  []string
	rows     [][]string
}

// NewTableBuilder creates a new TableBuilder instance
func NewTableBuilder(products []eoldate.Product) *TableBuilder {
	tb := &TableBuilder{products: products}
	tb.determineHeaders()
	tb.buildRows()
	return tb
}

// determineHeaders identifies all unique keys across all products
func (tb *TableBuilder) determineHeaders() {
	headerSet := make(map[string]bool)
	for _, product := range tb.products {
		v := reflect.ValueOf(product)
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if field.Name != "AdditionalFields" {
				tag := field.Tag.Get("json")
				if tag != "" && tag != "-" {
					headerName := strings.Split(tag, ",")[0]
					if !isEmptyValue(v.Field(i).Interface()) {
						headerSet[headerName] = true
					}
				}
			}
		}
	}

	tb.headers = make([]string, 0, len(headerSet))
	for header := range headerSet {
		if header != "" {
			tb.headers = append(tb.headers, header)
		}
	}
	sort.Strings(tb.headers)
}

// isEmptyValue checks if a value is considered empty
func isEmptyValue(v interface{}) bool {
	if v == nil {
		return true
	}
	switch value := v.(type) {
	case string:
		return value == ""
	case bool:
		return !value
	case int, int8, int16, int32, int64:
		return value == 0
	case float32, float64:
		return value == 0
	case []interface{}:
		return len(value) == 0
	case map[string]interface{}:
		return len(value) == 0
	case *float64:
		return value == nil
	default:
		return false
	}
}

// buildRows constructs the rows for the table
func (tb *TableBuilder) buildRows() {
	for _, product := range tb.products {
		row := make([]string, len(tb.headers))
		v := reflect.ValueOf(product)
		for i, header := range tb.headers {
			value := ""
			field := v.FieldByNameFunc(func(n string) bool {
				f, _ := v.Type().FieldByName(n)
				return strings.EqualFold(strings.Split(f.Tag.Get("json"), ",")[0], header)
			})
			if field.IsValid() {
				value = tb.formatValue(field.Interface())
			}
			if value == "" || value == "N/A" {
				if val, ok := product.AdditionalFields[header]; ok {
					value = tb.formatValue(val)
				}
			}
			row[i] = value
		}
		tb.rows = append(tb.rows, row)
	}
}

// formatValue converts an interface{} value to a string representation
func (tb *TableBuilder) formatValue(v interface{}) string {
	if v == nil {
		return "N/A"
	}

	switch value := v.(type) {
	case string:
		return value
	case float64:
		if value == float64(int64(value)) {
			return fmt.Sprintf("%.0f", value)
		}
		return fmt.Sprintf("%.2f", value)
	case *float64:
		if value == nil {
			return "N/A"
		}
		if *value == float64(int64(*value)) {
			return fmt.Sprintf("%.0f", *value)
		}
		return fmt.Sprintf("%.2f", *value)
	case bool:
		return fmt.Sprintf("%t", value)
	case time.Time:
		return value.Format("2006-01-02")
	case interface{}:
		return fmt.Sprintf("%v", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// Render creates and renders the table
func (tb *TableBuilder) Render() string {
	var buf strings.Builder
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(tb.headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	// Set header colors
	headerColors := make([]tablewriter.Colors, len(tb.headers))
	for i := range headerColors {
		headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor, tablewriter.BgBlackColor}
	}
	table.SetHeaderColor(headerColors...)

	for _, row := range tb.rows {
		colors := tb.colorizeRow(row)
		table.Rich(row, colors)
	}

	table.Render()
	return buf.String()
}

// colorizeRow applies color to specific columns based on their values
func (tb *TableBuilder) colorizeRow(row []string) []tablewriter.Colors {
	colors := make([]tablewriter.Colors, len(row))
	for i, header := range tb.headers {
		switch strings.ToLower(header) {
		case "eol", "support":
			colors[i] = tb.getDateColor(row[i])
		}
	}
	return colors
}

// getDateColor returns the appropriate color based on the date value
func (tb *TableBuilder) getDateColor(dateStr string) tablewriter.Colors {
	date, err := tb.parseDate(dateStr)
	if err != nil {
		return tablewriter.Colors{}
	}

	if date.Before(time.Now()) {
		return tablewriter.Colors{tablewriter.FgRedColor}
	}
	return tablewriter.Colors{tablewriter.FgGreenColor}
}

// parseDate attempts to parse a date string in various formats
func (tb *TableBuilder) parseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"2006-01",
		"2006",
	}

	for _, format := range formats {
		if date, err := time.Parse(format, dateStr); err == nil {
			return date, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}
