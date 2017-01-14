package main

import (
	//x "fmt"
	"errors"
	"strings"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/header"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
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
				return r_path, errors.New("Attempt to include duplcate ceh")
			}

			//d mooi_log("PATH: ", i, r_path)

			done, _ = utils.FileExists(r_path)
		}

		if done == false {
			err = errors.New("Include file not found: " + i_filename)
		}

		return
	}

	doProcess = func(i_lines chan readers.Line) (data_until bool, err error) {

		for kv := range readers.EachKeyVal(i_lines) {

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
				nestedLevel--
			} else {
				r_header.Add_kv(&kv)
				data_until = strings.EqualFold("DATA_UNTIL", kv.Key)
			}

			ix++
		}

		return
	}

	data_until, r_err := doProcess(i_lines_in)

	if r_err != nil {
		return
	}


	if data_until == false {
		r_err = errors.New("Error: data_until = false")
		return
	}

    r_header.Checks()

	diag.Info("Header Lines read:", ix)

	return
}
