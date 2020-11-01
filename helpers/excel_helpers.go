package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"github.com/tealeg/xlsx"
	"strings"
	"math"
	"strconv"
	"path/filepath"
	"archive/zip"
)
func ZipFolder(targetPath string, savePath string, fileName string) error {
	os.MkdirAll(savePath,0777)
	zipfile, err := os.Create(savePath+"/"+fileName)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(targetPath)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(targetPath)
	}

	filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if targetPath==path {
			return nil
		}
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = strings.TrimPrefix(path, targetPath)
		}

		if info.IsDir() {
			header.Name += "/"			
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func ReadCsvFile(filePath string) [][]string {
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	parsedData := make([][]string, 0, 0)


	for _, row := range result {

			var singleRow = make([]string,0,0)
			for _, col := range row {
				singleRow = append(singleRow,col)
			}
			if len(singleRow) > 0 {

				parsedData = append(parsedData, singleRow)
			}
		
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData
}

func ReadXlsxFile(filePath string) [][]string {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file")
	}

	parsedData := make([][]string, 0, 0)
	
	// sheet
	for _, sheet := range xlFile.Sheets {
		// rows
		for _, row := range sheet.Rows {

			// column
			var singleRow = make([]string,0,0)

			for _, cell := range row.Cells {
					text := cell.String()
					singleRow = append(singleRow,text)
			}
			if len(singleRow) > 0 {
				parsedData = append(parsedData, singleRow)
			}
		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))



	return parsedData
}
func SaveCSV(savePath string,fileName string, parsedData [][]string) {
	os.MkdirAll(savePath,0777)
	file, _ := os.Create(savePath+"/"+fileName)
	
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range parsedData {
       writer.Write(value)
    }
}
func SaveExcel(savePath string,fileName string, parsedData [][]string) {
	os.MkdirAll(savePath,0777)
	var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet")
    if err != nil {
        fmt.Printf(err.Error())
	}
	for _, row_value := range parsedData {
		row = sheet.AddRow()
		for _, cell_value := range row_value {
			cell = row.AddCell()
    		cell.Value = cell_value
		}
	}
   
    err = file.Save(savePath+"/"+fileName)
    if err != nil {
        fmt.Printf(err.Error())
    }
}
func ExcelCsvParser(blobPath string, blobExtension string) (parsedData [][]string) {
	fmt.Println("---------------> We are in product.go")
	if blobExtension == ".csv" {
		fmt.Println("-------We are parsing an csv file.-------------")
		parsedData := ReadCsvFile(blobPath)
		return parsedData

	} else if blobExtension == ".xlsx" {
		fmt.Println("----------------We are parsing an xlsx file.---------------")
		parsedData := ReadXlsxFile(blobPath)
		return parsedData
	}
	return parsedData
}

func SplitExcel(filePath string,records int, savePath string,keep_header int) (url string) {

	i := strings.LastIndex(filePath, ".")
	i2 := strings.LastIndex(filePath, "/")+1

	extension :=filePath[i:len(filePath)]
	fileName:= filePath[i2:i]

	parsedData := ExcelCsvParser(filePath,extension)
	header := parsedData[0]
	if(keep_header==1) {
		parsedData = parsedData[1:len(parsedData)]
	}
	files := int(math.Ceil( float64(len(parsedData))/ float64(records)))
	for i := 1; i <= files; i++ {
		from := (i-1) * records
		to := i * records
		if to>len(parsedData){
			to = len(parsedData)
		}
		data_save := parsedData[from:to]
		if(keep_header==1) {
			data_save = append([][]string{header}, data_save...)
		}
		if extension==".csv" {
			SaveCSV(savePath,fileName+"_"+strconv.Itoa(i)+".csv",data_save) 
		} else if extension==".xlsx"{
			SaveExcel(savePath,fileName+"_"+strconv.Itoa(i)+".xlsx",data_save) 
		}
	}
	
	return extension

}
func Test() {

}