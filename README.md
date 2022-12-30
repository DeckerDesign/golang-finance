golang-finance

Welcome to golang-finance, a Go program for managing your budget. This program allows you to enter your expenses and amounts, and then updates a budget spreadsheet (budget.xlsx) with this information.

Getting Started

To use golang-finance, you will need to have Go and the github.com/tealeg/xlsx library installed on your system. You can install Go by following the instructions here. To install the xlsx library, run the following command:

```Copy codego get github.com/tealeg/xlsx```

Once you have Go and the xlsx library installed, you can clone this repository and run the program:

```Copy codegit clone https://github.com/DeckerDesign/golang-finance.git```
```cd golang-finance```
```go run main.go```
The program will prompt you to enter the names of your expenses and the amounts. When you are done, enter done to save the changes to the budget.xlsx file.

Duplicating the Data

To duplicate the data in the budget spreadsheet, you can run the duplicate.go program:

```Copy code go run duplicate.go```
This program will create a new sheet in the budget.xlsx file called "Table", and will paste the data from the original sheet 12 times, going down. This will allow you to easily see a summary of your budget data.

Viewing the Data

To view the data in the budget spreadsheet, you can use a program such as Microsoft Excel or Google Sheets. You can then create a table or pivot table to analyze and drill down into the data.

I hope you find golang-finance useful for managing your budget! Let me know if you have any questions or need further assistance.
