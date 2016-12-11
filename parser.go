package main

import (
	"errors"
	"strings"
    "fmt"
)

func ReadCefHeader(args *CefArgs) (r_header CefHeaderData, r_lines chan Line, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0

    r_header = *NewCefHeaderData()
	l_path := *args.m_cefpath

	// forward decl
	var doProcess func(i_lines chan Line) (data_until bool, err error)

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

			mooi_log("PATH: ", i, r_path)

			done, _ = fileExists(r_path)
		}

		if done == false {
			err = errors.New("Include file not found: " + i_filename)
		}

		return
	}

	doProcess = func(i_lines chan Line) (data_until bool, err error) {

		for kv := range eachKeyVal(i_lines) {

			if strings.EqualFold("include", kv.key) == true {
				v := strings.Trim(kv.val[0], `" `)

				ceh_path, err := getIncludePath(v)
				if err != nil {
					return data_until, err
				}

				includedMap[ceh_path] = true
				nestedLevel++

				l_lines := EachLine(ceh_path)				
				if _, err = doProcess(l_lines); err != nil {
					return data_until, err
				}
				nestedLevel--
			} else {
				r_header.add_kv(&kv)
				data_until = strings.EqualFold("DATA_UNTIL", kv.key)
			}

			ix++
		}

		return
	}

	r_lines = EachLine(l_path)
	data_until, r_err := doProcess(r_lines)
	if r_err != nil {
		return
	}

	if data_until == false {
		r_err = errors.New("Error: data_until = false")
		return
	}

    r_header.checks()

    //-r_header.dumpAttrs()
	println("Lines read -> ", ix)

	//- r_header.m_data.dump()
    
    println("//x (cef_header.go) r_header.dump()")
	//x r_header.dump()

	return
}

//x type Line struct {
//x     tag         string     // typically filename
//x     ln          int        // line number
//x     line        string     // line contents
//x }


// TODO 2016/2017
	- User Header Data to get access to info on Vars.

	- Multi-Line Data Reader
		- Check comments
		- Check number of elements
		- Check eol marker
		- Check type of each cell
		- Check Time/Date stamp 


func ReadCef(args *CefArgs) (r_err error) {

	_, l_lines, err := ReadCefHeader(args)
	error_check(err, "Error parsing header")

	// dev data line reader
	ix := 0
	for l_line := range l_lines {
		fmt.Println("line:", ix, l_line)

		if ix > 3 {
			break
		}

		ix++
	}

	return
}

