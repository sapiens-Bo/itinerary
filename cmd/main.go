package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	const sheet = "TDSheet"
	f, err := excelize.OpenFile("book.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := f.GetCellValue(sheet, "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Cell B2: %s\n", cell)

	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell, "\t")
		}
		fmt.Println()
	}
}
