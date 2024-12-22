package main

import (
	"fmt"

	"github.com/sapiens-Bo/itinerary/internal/excel"
)

func main() {
	/*
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
		fmt.Printf("Cell B2: %s\n-------------------\n", cell)

		stID, err := f.GetCellStyle(sheet, "A1")
		if err != nil {
			fmt.Println(err)
			return
		}

		style := f.Styles.CellXfs.Xf[stID]
		fontColor := style.FillID

		fmt.Printf("Style cell B2: %v\n", *fontColor)

		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, row := range rows {
			for _, colCell := range row {
				fmt.Printf("\n%T\n", colCell)
				fmt.Println(colCell, "\t")
			}
			fmt.Println()
		}
	*/
	fmt.Println("===Tutor===")
	excel.RowStuff()
}
