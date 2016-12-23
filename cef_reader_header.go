package main

import (
	"errors"
	"strings"
    "github.com/caa-dev-apps/cefmdd_v1/readers"
    "github.com/caa-dev-apps/cefmdd_v1/utils"
)

func ReadHeader(args *CefArgs,
			    i_lines_in chan readers.Line) (r_header CefHeaderData, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0

    r_header = *NewCefHeaderData()

	// forward decl
	var doProcess func(i_lines chan readers.Line) (data_until bool, err error)

	getIncludePath := func(i_filename string) (r_path string, err error) {
		done := false

		if nestedLevel >= 8 {
			return r_path, errors.New("Include nested files limit reached (8 deep)")
		}

		for i := 0; i < len(args.m_includes) && done == false; i++ {

			r_path = args.m_includes[i] + `/` + i_filename

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
				r_header.add_kv(&kv)
				data_until = strings.EqualFold("DATA_UNTIL", kv.Key)
			}

			ix++
		}

		return
	}

	//x r_lines = EachLine(l_path)
	//x data_until, r_err := doProcess(r_lines)
	data_until, r_err := doProcess(i_lines_in)


	if r_err != nil {
		return
	}



	if data_until == false {
		r_err = errors.New("Error: data_until = false")
		return
	}

    r_header.checks()

    //-r_header.dumpAttrs()
	println("Header Lines read -> ", ix)

	//- r_header.m_data.dump()
    
    println("//x (cef_header.go) r_header.dump()")
	//x r_header.dump()

	return
}
