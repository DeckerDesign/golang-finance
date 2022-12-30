package main

import (
	"bufio"
	"fmt"
	"os"

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

	// Set up a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Set up counters for the current row and column
	row := 6
	col := 2

	// Print a prompt for the user to enter the name of the expenses
	fmt.Println("Please enter the name of the expenses. When done, enter 'done'.")

	// Read user input for the names of the expenses
	for {
		// Print a prompt and scan for user input
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// If the user enters "done", break out of the loop
		if input == "done" {
			break
		}

		// Otherwise, add the input to the current cell in the sheet
		cell := sheet.Cell(row, col)
		cell.Value = input

		// Increment the row and move to the next cell
		row++
	}

	// Reset the row counter
	row = 6

	// Print a prompt for the user to enter the amounts
	fmt.Println("Please enter the amounts. When done, enter 'done'.")

	// Read user input for the amounts
	for {
		// Print a prompt and scan for user input
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// If the user enters "done", break out of the loop
		if input == "done" {
			break
		}

		// Otherwise, add the input to the current cell in the sheet
		cell := sheet.Cell(row, col+1)
		cell.Value = input

		// Increment the row and move to the next cell
		row++
	}

	// Save the changes to the file
	err = xlFile.Save(excelFileName)
	if err != nil {
		fmt.Printf("Error saving file: %s\n", err)
		return
	}
}
