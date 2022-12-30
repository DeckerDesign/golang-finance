package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	// Open the budget.xlsx file
	excelFileName := "budget.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}

	// Get the first sheet
	sheet := xlFile.Sheets[0]

	// Create a new sheet for the table
	newSheet, err := xlFile.AddSheet("Table")
	if err != nil {
		fmt.Printf("Error adding sheet: %s\n", err)
		return
	}

	// Copy the data from the original sheet to the new sheet
	for r := 0; r < sheet.MaxRow; r++ {
		newRow := newSheet.AddRow()
		for c := 0; c < sheet.MaxCol; c++ {
			cell := sheet.Cell(r, c)
			newCell := newRow.AddCell()
			newCell.Value = cell.Value
		}
	}

	// Paste the data 12 times, going down
	for i := 0; i < 12; i++ {
		for r := 0; r < sheet.MaxRow; r++ {
			newRow := newSheet.AddRow()
			for c := 0; c < sheet.MaxCol; c++ {
				cell := sheet.Cell(r, c)
				newCell := newRow.AddCell()
				newCell.Value = cell.Value
			}
		}
	}

	// Save the changes to the file
	err = xlFile.Save(excelFileName)
	if err != nil {
		fmt.Printf("Error saving file: %s\n", err)
		return
	}
}
