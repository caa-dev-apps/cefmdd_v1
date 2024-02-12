package main

import (
	"errors"
	"strings"
	"fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/header"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/rules"
)

var (
    s_mdd_data = *rules.NewMddData()
)


func ReadHeader(args *utils.CefArgs,
			    i_lines_in chan readers.Line) (r_header header.CefHeaderData, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0

    r_header = *header.NewCefHeaderData()

	// forward decl
	var doProcess func(i_lines chan readers.Line) (data_until bool, err error)

	getIncludePath := func(i_filename string) (r_path string, err error) {
		done := false

		if nestedLevel >= 8 {
			return r_path, errors.New("Include nested files limit reached (8 deep)")
		}

		l_includes := args.GetIncludes()
		for i := 0; i < len(l_includes) && done == false; i++ {

			r_path = l_includes[i] + `/` + i_filename

			_, included := includedMap[r_path]
			if included == true {
				diag.Error("Attempt to include duplicate ceh", i_filename)
				return r_path, errors.New("Attempt to include duplicate ceh")
			}

			done, _ = utils.FileExists(r_path)
		}

		if done == false {
			diag.Error("Include file not found: ", i_filename)
			err = errors.New("Include file not found: " + i_filename)
		}

		return
	}

	doProcess = func(i_lines chan readers.Line) (data_until bool, err error) {

		for kv := range readers.EachKeyVal(i_lines) {
//202001 smcc added begin
		count:=strings.Count(fmt.Sprintf("%s",kv.Line),"\"")

		if (count%2 == 0){
			//This string either has one or none quotes
		}else{
			err = errors.New(fmt.Sprintf("Value mismatched quotes %s", kv.Line))
			return
		}
//202001 smcc added end
			if strings.EqualFold("include", kv.Key) == true {
				v := strings.Trim(kv.Val[0], `" `)

				ceh_path, err := getIncludePath(v)
				if err != nil {
					return data_until, err
				}

				includedMap[ceh_path] = true
				nestedLevel++

				l_lines := readers.EachLine(ceh_path)				
				if _, err = doProcess(l_lines); err != nil {
					return data_until, err
				}

				diag.Trace("File Close", ceh_path)

				nestedLevel--
			} else {

				err2 := r_header.Add_kv(&kv)
				if err2 != nil {
					return false, err2
				}

				data_until = strings.EqualFold("DATA_UNTIL", kv.Key)
			}

//		fmt.Println(kv)
			ix++
		}

	return false,err //2019
	}

	data_until, r_err := doProcess(i_lines_in)

	if r_err != nil {
		return
	}
	if data_until == false {
		r_err = errors.New("Error: data_until = false")
	}

	r_err = r_header.CheckMin("m")  // Check mission level
	  if r_err == nil{
	    r_err = r_header.Checks()  //smcc
	}

	diag.Info("Header Lines read:", ix)

	return  
}



