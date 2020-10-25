package helpers

import (
	"bufio"
	"container/list"
	"encoding/csv"
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
	// "net/http"
	"os"
	// "path/filepath"
	"github.com/tealeg/xlsx"
	//"strings"
)

func ReadCsvFile(filePath string) []map[string]interface{} {
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	parsedData := make([]map[string]interface{}, 0, 0)
	header_name := result[0]

	for row_counter, row := range result {

		if row_counter != 0 {
			var singleMap = make(map[string]interface{})
			for col_counter, col := range row {
				singleMap[header_name[col_counter]] = col
			}
			if len(singleMap) > 0 {

				parsedData = append(parsedData, singleMap)
			}
		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData

}

func ReadXlsxFile(filePath string) []map[string]interface{} {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file")
	}

	parsedData := make([]map[string]interface{}, 0, 0)
	header_name := list.New()
	// sheet
	for _, sheet := range xlFile.Sheets {
		// rows
		for row_counter, row := range sheet.Rows {

			// column
			header_iterator := header_name.Front()
			var singleMap = make(map[string]interface{})

			for _, cell := range row.Cells {
				if row_counter == 0 {
					text := cell.String()
					header_name.PushBack(text)
				} else {
					text := cell.String()
					singleMap[header_iterator.Value.(string)] = text
					header_iterator = header_iterator.Next()
				}
			}
			if row_counter != 0 && len(singleMap) > 0 {

				parsedData = append(parsedData, singleMap)
			}

		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData
}

func ExcelCsvParser(blobPath string, blobExtension string) (parsedData []map[string]interface{}) {
	fmt.Println("---------------> We are in product.go")
	if blobExtension == ".csv" {
		fmt.Println("-------We are parsing an csv file.-------------")
		parsedData := ReadCsvFile(blobPath)
		fmt.Printf("Type:%T\n", parsedData)
		return parsedData

	} else if blobExtension == ".xlsx" {
		fmt.Println("----------------We are parsing an xlsx file.---------------")
		parsedData := ReadXlsxFile(blobPath)
		return parsedData
	}
	return parsedData
}
func Test() (){

}