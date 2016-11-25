package gsheet

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

type Wrapper struct {
	spreadsheetID string
	srv           *sheets.Service
}

func NewWrapper(spreadsheetID string, srv *sheets.Service) (s *Wrapper) {
	s = &Wrapper{}
	s.spreadsheetID = spreadsheetID
	s.srv = srv
	return
}

func (s *Wrapper) BatchClear(clearRange string) (err error) {

	req := &sheets.BatchClearValuesRequest{}
	req.Ranges = []string{clearRange}

	fmt.Println("Clearing target sheet " + s.spreadsheetID)
	_, err = s.srv.Spreadsheets.Values.BatchClear(s.spreadsheetID, req).Do()
	return
}

func (s *Wrapper) Update(writeRange string, sheetRow SheetDataInput) (
	err error) {

	values := &sheets.ValueRange{}
	values.Range = writeRange
	values.Values = sheetRow.GetSheetData()
	values.MajorDimension = "ROWS"
	valueInputOpt := &Option{
		key:   "valueInputOption",
		value: "USER_ENTERED",
	}

	fmt.Println("Writing to target sheet " + s.spreadsheetID)
	_, err = s.srv.Spreadsheets.Values.Update(s.spreadsheetID, writeRange,
		values).Do(valueInputOpt)
	if err != nil {
		return
	}

	fmt.Println("Success updated sheet")

	return
}
