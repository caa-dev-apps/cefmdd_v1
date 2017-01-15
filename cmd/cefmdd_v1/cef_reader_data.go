package main

import (
	"errors"
    "fmt"
    "time"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/data"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/header"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

//.     type Line struct {
//.         tag         string     // typically filename
//.         ln          int        // line number
//.         line        string     // line contents
//.     }
//.     
//.      TODO 2016/2017
//.     	- User Header Data to get access to info on Vars.
//.     
//.     	- Multi-Line Data Reader
//.     		- Check comments
//.     		- Check number of elements
//.     		- Check eol marker
//.     		- Check type of each cell
//.    		- Check Time/Date stamp 
//.     
//.     type Cell struct {
//.     	Variable_ix 	int
//.     	Cell_ix 		int
//.     	VariableType 	string
//.     	VariableParser  func(string) (error)
//.     }
//. 	type CefRecord struct {
//. 	    Err error
//. 	    Tokens []string
//. 	    //x Line string       
//. 	    Line Line       
//. 	}
//. 	
//. 	func DataRecords(i_lines chan Line, i_len int) chan CefRecord {

func ReadData(i_header header.CefHeaderData,
			  i_lines chan readers.Line) (err error) {


	l_cellTypes, err := data.RecordCellTypes(i_header)
	if err != nil {
		diag.Error("Creating Cell Types Array for Data checks")
		return
	}

	var t0 time.Time

	r_ix := 0
	for l_record := range readers.DataRecords(i_lines, len(l_cellTypes))  {

		if l_record.Err != nil {
			err = l_record.Err
			diag.Errorf("Reading data record ", "%#v", l_record.Line)
			return
		}

		for ix, ct := range l_cellTypes {
			rd := l_record.Tokens[ix]
			err = ct.VariableParser(rd) 
			if err != nil {

				diag.Errorf("In data record, incorrect data type\n",
							"%s\n" +
							"Error Token %s\n" +
							"Tokens %#v\n" +
							"Line %#v\n",
							ct,
							rd,
							l_record.Tokens,
							l_record.Line)
				return err
			}
		}

		t1, _ := utils.Iso_time(l_record.Tokens[0])

		if(t1.Sub(t0) < 0) {
			return errors.New(fmt.Sprintf("Error in Record Date/Time stamp predating previous record value t0(%#v) t1(%#v) %#v", t0, t1, l_record.Line))
		}

		t0 = t1

		r_ix++
		if r_ix >= 100 {
			break
		}

	}

	diag.Info("Data Records Checked: ", r_ix)

	return
}