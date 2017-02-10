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
//. 	    //- Line string       
//. 	    Line Line       
//. 	}
//. 	
//. 	func DataRecords(i_lines chan Line, i_len int) chan CefRecord {

func ReadData(i_header header.CefHeaderData,
			  i_lines chan readers.Line) (err error) {

	var l_data_until string 
	var l_end_of_record_marker string 

    du, p := i_header.Attrs().Map()["DATA_UNTIL"];
    if p == true {
    	if du[0] != "EOFx" {
    		l_data_until = du[0]
    	}
    }

	eorms, p := i_header.Attrs().Map()["END_OF_RECORD_MARKER"]
	if p == true {
		if len(eorms) == 1 {
			l_end_of_record_marker = utils.Trim_quoted_string(eorms[0])
		} else {
			diag.Error("Invalid END_OF_RECORD_MARKER", eorms)
			return
		}
	}


	l_cellTypes, err := data.RecordCellTypes(i_header)
	if err != nil {
		diag.Error("Creating Cell Types Array for Data checks")
		return
	}

	is_cell_0_ISO_TIME := l_cellTypes[0].VariableType == "ISO_TIME" 

	var t0 time.Time

	l_max_records := utils.GetCefArgs().GetMaxRecords()
	r_ix := 0
	for l_record := range readers.DataRecords(i_lines, len(l_cellTypes), l_data_until, l_end_of_record_marker)  {

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

		if is_cell_0_ISO_TIME == true {
			t1, _ := utils.Iso_time(l_record.Tokens[0])

			if(t1.Sub(t0) < 0) {
				return errors.New(fmt.Sprintf("Error in Record Date/Time stamp predating previous record value t0(%#v) t1(%#v) %#v", t0, t1, l_record.Line))
			}

			t0 = t1
		}

		// max_records = -1 for all
		r_ix++;
		if l_max_records > 0 && r_ix >= l_max_records {
			break
		}

	}

	diag.Info("Data Records Checked: ", r_ix)

	return
}