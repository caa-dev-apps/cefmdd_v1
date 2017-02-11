package data

import (
    "fmt"
    "errors"
    "strconv"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/header"
 	"github.com/caa-dev-apps/cefmdd_v1/pkg/utils"    
//x 	"github.com/caa-dev-apps/cefmdd_v1/pkg/diag"    
)

type RecordCellType struct {
	Variable_ix 	int
	Cell_ix 		int
	VariableName    string
	VariableType 	string
	VariableParser  func(string) (error)
}

func(ct RecordCellType) String() (s string) {

	return fmt.Sprintf("Variable Name: %s, " +
						"Variable Index: %d, " +
						"Variable Type: %s, " +
						"Cell Index: %d ",
						ct.VariableName,
						ct.Variable_ix,
						ct.VariableType,
						ct.Cell_ix)
}

//x func valueTypeParserFunc(vt string) (f func(string) (error), err error) {
//x 
//x 	switch vt {
//x 		case "CHAR" : 				f = utils.String_parser
//x 		case "INT" :				f = utils.Integer_parser
//x 		case "FLOAT" :				f = utils.Numerical_parser
//x 		case "DOUBLE" :				f = utils.Numerical_parser
//x 		case "ISO_TIME" :			f = utils.Iso_time_parser
//x 		case "ISO_TIME_RANGE" :		f = utils.Iso_time_range_parser
//x 		default :
//x 			err = errors.New("Unknown VALUE_TYPE : " + vt)
//x 	}
//x 
//x 	return
//x }

//x func SizesProduct(sizes []string) (product int64, err error) {
//x 
//x 	for i, v := range sizes {
//x 	    n, err1 := strconv.ParseInt(v, 10, 64)
//x 	    if err1 != nil {
//x 	        return 0, err1
//x 	    }
//x 
//x 	    if i == 0 {
//x 	    	product = n 
//x 	    } else {
//x 	    	product *= n
//x 	    }
//x 	}
//x 
//x 	if product == 0 {
//x     	err = errors.New(fmt.Sprintf(`SIZES malformed %#v`, sizes))
//x 	}
//x     
//x     return
//x }

func RecordCellTypes(i_header header.CefHeaderData) (r_cells []RecordCellType, err error) {
	vs := i_header.Vars().List()

//x	var debug_sz_tot int64 

	c_ix := 0
	for ix, v := range vs {

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
			err = errors.New(fmt.Sprintf(`error SIZES missing from Variable %#v`, vn))
			return
		}

		sp, err1 := utils.SizesProduct(sizes)  
		if err1 != nil {
			err = errors.New(fmt.Sprintf(`error SIZES malformed in Variable %#v`, vn))
			return
		}

//x		diag.Info("SIZES", sp, sizes)
//x		debug_sz_tot += sp

		vt, p := vmap["VALUE_TYPE"]
		if p == false {
			err = errors.New(fmt.Sprintf(`error VALUE_TYPE missing from Variable %v`, vn))
			return
		}

		f, err2 := utils.ValueTypeParserFunc(vt[0])
		if err2 != nil {
			err = errors.New(fmt.Sprintf(`error VALUE_TYPE unknown type : %v  -  variable : %v`, vt, vn))
			return
		} 

		for sp_ix := int64(0); sp_ix < sp; sp_ix++ {

			l_cell := &RecordCellType{Variable_ix: ix,
							Cell_ix: c_ix,
							VariableName: vn[0],
							VariableType: vt[0],
							VariableParser: f,
						}

			r_cells = append(r_cells, *l_cell)
			
			c_ix++
        }

//x		diag.Info("SIZES-TOTAL", debug_sz_tot)

    }

	amap := i_header.Attrs().Map()
	eorms, p := amap["END_OF_RECORD_MARKER"]
	if p == true && len(eorms) == 1 {
		e := utils.Trim_quoted_string(eorms[0])

		l_cell := &RecordCellType{Variable_ix: -1,
						Cell_ix: c_ix,
						VariableName: "END_OF_RECORD_MARKER",
						VariableType: "string",
						VariableParser: utils.String_matcher_test(e),
					}

		r_cells = append(r_cells, *l_cell)
	}

	return
}

