package excel

import (
	"errors"
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

var bgcolorCell = "FF000000"

func cellVisitor(c *xlsx.Cell) error {
	style := c.GetStyle()
	strBG := style.Fill.BgColor
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else if value != "" {
		fmt.Printf("Cell value: %s BackColor: %s\n", value, strBG)
	}

	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func RowStuff() {
	filename := "book.xlsx"
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["TDSheet"]
	if !ok {
		panic(errors.New("sheet not found"))
	}

	fmt.Printf("Max row is: %d\n", sh.MaxRow)
	sh.ForEachRow(rowVisitor)
}
