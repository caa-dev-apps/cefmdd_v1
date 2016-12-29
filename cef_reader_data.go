package main

import (
    "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/header"
    "github.com/caa-dev-apps/cefmdd_v1/readers"
//x 	"github.com/caa-dev-apps/cefmdd_v1/utils"    
)


//x type Line struct {
//x     tag         string     // typically filename
//x     ln          int        // line number
//x     line        string     // line contents
//x }


//!  TODO 2016/2017
//! 	- User Header Data to get access to info on Vars.
//! 
//! 	- Multi-Line Data Reader
//! 		- Check comments
//! 		- Check number of elements
//! 		- Check eol marker
//! 		- Check type of each cell
//! 		- Check Time/Date stamp 


type Cell struct {
	Variable_ix 	int
	Cell_ix 		int
	Variable_type 	int
}

func NewCellsArray(i_header header.CefHeaderData) (r_cells []Cell) {

	vs := i_header.Vars().List()

	for ix, v := range vs {
		fmt.Println(ix)		

		for k1, v1 := range v.Map() {
			fmt.Println(k1, "\t", v1)	

			// TODO move to []struct {
			// 	Variable_Name string
			// 	Value_Type string
			// 	Sizes_total n0 x n1 x n2...
			// }
			// 
		}

		//x fmt.Println(v)		
		fmt.Println("\n\n")		
	}

	//x fmt.Println(vs)

	return
}



func ReadData(i_header header.CefHeaderData,
			  i_lines chan readers.Line) (r_err error) {

	// dev data line reader
	ix := 0
	for l_line := range i_lines {
		fmt.Println("line:", ix, l_line)

		if ix > 3 {
			break
		}

		ix++
	}

	_ = NewCellsArray(i_header)

	return
}

