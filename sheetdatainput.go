package gsheet

type SheetDataInput interface {
	GetSheetData() [][]interface{}
}
