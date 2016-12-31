package data

import (
    "fmt"
    "errors"
    "strconv"
    "github.com/caa-dev-apps/cefmdd_v1/header"
 	"github.com/caa-dev-apps/cefmdd_v1/utils"    
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

type DataCellType struct {
	Variable_ix 	int
	Cell_ix 		int
	VariableName    string
	VariableType 	string
	VariableParser  func(string) (error)
}

//		CATDESC 	 						["Cluster C3, theta look direction index for D2"]			
//		variable_name 	 		x			[D2_theta_index__C3_CP_EDI_QZC]						
//		LABLAXIS 	 						["i_theta_D2"]						
//		VALUE_TYPE 	 			x			[INT]			
//		SIGNIFICANT_DIGITS 	 				[3]
//		SIZES 	 				x			[1]												
//		FIELDNAM 	 						["Cluster C3, theta look direction index for D2"]
//		DEPEND_0 	 						[time_tags__C3_CP_EDI_QZC]
//		PARAMETER_TYPE 	 					["Support_Data"]
//		PROPERTY 	 						["Status"]
//		SI_CONVERSION 	 					["1>unitless"]
//		FILLVAL 	 						[255]
//		FRAME 	 							["scalar>na"]
//		ENTITY 	 							["Instrument"]
//		UNITS 	 							["unitless"]
//		DATA 					x 			["unitless"]

//		[CHAR]               
//		[INT]                
//		[FLOAT]              
//		[ISO_TIME]           
//		[ISO_TIME_RANGE]     

//x func Integer_parser(v string) (err error) { 
//x func Iso_time_parser(v string) (err error) { 
//x func Iso_time_range_parser(v string) (err error) { 
//x func Numerical_parser(v string) (err error) { 
//x func String_parser(v string) (err error) { 

func valueTypeParserFunc(vt string) (f func(string) (error), err error) {

	switch vt {
		case "CHAR" : 				f = utils.String_parser
		case "INT" :				f = utils.Integer_parser
		case "FLOAT" :				f = utils.Numerical_parser
		case "ISO_TIME" :			f = utils.Iso_time_parser
		case "ISO_TIME_RANGE" :		f = utils.Iso_time_range_parser
		default :
			err = errors.New("Unknown VALUE_TYPE : " + vt)
	}

	return
}

func sizesProduct(sizes []string) (product int64, err error) {

	for i, v := range sizes {
	    n, err1 := strconv.ParseInt(v, 10, 64)
	    if err1 != nil {
	        return 0, err1
	    }

	    if i == 0 {
	    	product = n 
	    } else {
	    	product *= n
	    }
	}

	if product == 0 {
    	err = errors.New(fmt.Sprintf(`SIZES malformed %#v`, sizes))
	}
    
    return
}

func NewDataCellTypesArray(i_header header.CefHeaderData) (r_cells []DataCellType, err error) {
	vs := i_header.Vars().List()

	c_ix := 0
	for ix, v := range vs {
		//x fmt.Println(ix)		

		vmap := v.Map()
		if _, p := vmap["DATA"]; p == true {
			// DATA variable - is a constant
			continue
		}

		// not official - but should get added for convenience
		vn, p := vmap["variable_name"] 
		if p == false {
			err = errors.New(`error variable_name missing from Variable index: ` + strconv.Itoa(ix))
			return
		}

		// size = an array of numbers - as strings
		sizes, p := vmap["SIZES"]
		if p == false {
			//x err = errors.New(`error SIZES missing from Variable ` + vn)
			err = errors.New(fmt.Sprintf(`error SIZES missing from Variable %#v`, vn))
			return
		}

		sp, err1 := sizesProduct(sizes)  
		if err1 != nil {
			err = errors.New(fmt.Sprintf(`error SIZES malformed in Variable %#v`, vn))
			return
		}

		vt, p := vmap["VALUE_TYPE"]
		if p == false {
			err = errors.New(fmt.Sprintf(`error VALUE_TYPE missing from Variable %#v`, vn))
			return
		}

		f, err2 := valueTypeParserFunc(vt[0])
		if err2 != nil {
			err = errors.New(fmt.Sprintf(`error VALUE_TYPE unknown type : %#v  -  variable : %#v`, vt, vn))
			return
		} 

		for sp_ix := int64(0); sp_ix < sp; sp_ix++ {

			l_cell := &DataCellType{Variable_ix: ix,
							Cell_ix: c_ix,
							VariableName: vn[0],
							VariableType: vt[0],
							VariableParser: f,
						}

			r_cells = append(r_cells, *l_cell)
			
			c_ix++
        }
    }

	amap := i_header.Attrs().Map()
	eorms, p := amap["END_OF_RECORD_MARKER"]
	if p == true && len(eorms) == 1 {
		l_cell := &DataCellType{Variable_ix: -1,
						Cell_ix: c_ix,
						VariableName: "END_OF_RECORD_MARKER",
						VariableType: "string",
						VariableParser: utils.String_matcher_test(eorms[0]),
					}

		r_cells = append(r_cells, *l_cell)
	}

	return
}

