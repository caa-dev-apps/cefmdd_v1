package main

import (
    "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/diag"
    "github.com/caa-dev-apps/cefmdd_v1/data"
    "github.com/caa-dev-apps/cefmdd_v1/header"
    "github.com/caa-dev-apps/cefmdd_v1/readers"
)

//..    //x type Line struct {
//..    //x     tag         string     // typically filename
//..    //x     ln          int        // line number
//..    //x     line        string     // line contents
//..    //x }
//..    
//..    //!  TODO 2016/2017
//..    //! 	- User Header Data to get access to info on Vars.
//..    //! 
//..    //! 	- Multi-Line Data Reader
//..    //! 		- Check comments
//..    //! 		- Check number of elements
//..    //! 		- Check eol marker
//..    //! 		- Check type of each cell
//..    //! 		- Check Time/Date stamp 
//..    
//..    type Cell struct {
//..    	Variable_ix 	int
//..    	Cell_ix 		int
//..    	VariableType 	string
//..    	VariableParser  func(string) (error)
//..    }

// func rowTokeniser(i_lines chan string, i_len int) chan RowTokens 

func ReadData(i_header header.CefHeaderData,
			  i_lines chan readers.Line) (err error) {

	// dev data line reader
	//x ix := 0
	//x for l_line := range i_lines {
	//x 	fmt.Println("line:", ix, l_line)
//x 
//x 		if ix > 3 {
//x 			break
//x 		}
//x 
//x 		ix++
//x 	}

	l_cells, err := data.NewDataCellTypesArray(i_header)
	if err != nil {
		fmt.Println(diag.BoldRed("Error creating Cells Array for Data checks"))	
		return
	}

	for i, c := range l_cells {
		//x fmt.Println(i)
		fmt.Printf("%d \t %#v \n", i, c)
	}




//x type RowTokens struct {
//x     Err error
//x     Tokens []string
//x     //x Line string       // 
//x     Line Line       // 
//x }

	//x chan RowTokens
	ix := 0
	for l_tokens := range readers.RowTokeniser(i_lines, len(l_cells))  {

		if l_tokens.Err != nil {
			err = l_tokens.Err
			// DO SOME OUTPUT including FILE + LINE number
			return
		}

		fmt.Println(ix)
		//x fmt.Printf("%#v\n", l_tokens)
		fmt.Printf("Line: %#v\n", l_tokens.Line)
		fmt.Printf("Tokens: %#v\n", l_tokens.Tokens)
		fmt.Printf("\n")

		if ix > 3 {
			break
		}
		ix++

	}


	return
}

