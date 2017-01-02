package main

import (
	"errors"
    "fmt"
    "time"
    "github.com/caa-dev-apps/cefmdd_v1/diag"
    "github.com/caa-dev-apps/cefmdd_v1/data"
    "github.com/caa-dev-apps/cefmdd_v1/header"
    "github.com/caa-dev-apps/cefmdd_v1/readers"
    "github.com/caa-dev-apps/cefmdd_v1/utils"
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
		fmt.Println(diag.BoldRed("Error creating Cell Types Array for Data checks"))	
		return
	}

	var t0 time.Time

	r_ix := 0
	for l_record := range readers.DataRecords(i_lines, len(l_cellTypes))  {

		if l_record.Err != nil {
			err = l_record.Err
			fmt.Println(diag.BoldRed(fmt.Sprintf("Error reading data record %#v", l_record.Line)))
			return
		}

		for ix, ct := range l_cellTypes {
			rd := l_record.Tokens[ix]
			err = ct.VariableParser(rd) 
			if err != nil {

				fmt.Println(diag.BoldRed(fmt.Sprintf("Error in data record, incorrect data type\n" +
													 "%s\n" +
													 "Error Token %s\n" +
													 "Tokens %#v\n" +
													 "Line %#v\n",
													 ct,
													 rd,
													 l_record.Tokens,
													 l_record.Line)))
				return err
			}
		}

		t1, _ := utils.Iso_time(l_record.Tokens[0])
		//- fmt.Printf("t0 %#v\t  t1 %#v \n", t0, t1)

		if(t1.Sub(t0) < 0) {
			return errors.New(fmt.Sprintf("Error in Record Date/Time stamp predating previous record value t0(%#v) t1(%#v) %#v", t0, t1, l_record.Line))
		}

		t0 = t1

		if r_ix > 3 {
			break
		}
		r_ix++

	}

	return
}