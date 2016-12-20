package main

import (
    "fmt"
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


func ReadData(args *CefArgs,
			  i_lines chan Line) (r_err error) {

	// dev data line reader
	ix := 0
	for l_line := range i_lines {
		fmt.Println("line:", ix, l_line)

		if ix > 3 {
			break
		}

		ix++
	}

	return
}

