package main

import (
	"errors"
	"strings"
    "fmt"
)

func ReadCefHeader(args *CefArgs) (r_header CefHeaderData, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0

	//x r_header = CefHeaderData{m_state: ATTR, m_data: Parsed_Data{}}
    r_header = *NewCefHeaderData()
    
	l_path := *args.m_cefpath

	// forward decl
	var doProcess func(i_path string) (data_until bool, err error)

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

	doProcess = func(i_filepath string) (data_until bool, err error) {
		l_lines := EachLine(i_filepath)

		for kv := range eachKeyVal(l_lines) {

			if strings.EqualFold("include", kv.key) == true {
				v := strings.Trim(kv.val[0], `" `)

				ceh_path, err := getIncludePath(v)
				if err != nil {
					return data_until, err
				}

				includedMap[ceh_path] = true
				nestedLevel++

				if _, err = doProcess(ceh_path); err != nil {
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

	data_until, r_err := doProcess(l_path)
	if r_err != nil {
		return
	}

	if data_until == false {
		r_err = errors.New("Error: data_until = false")
		return
	}

    r_header.dumpAttrs()
    

    
    aks := []string {
        "DATA_UNTIL",
        "FILE_FORMAT_VERSION",
        "FILE_NAME",
    }
    
    for _, k := range aks {
        if v, err := r_header.getAttr(k); err == nil {
            fmt.Println(k, v)
        } else {
            fmt.Println(err)
        }
        
    }
    
    
    mks := []string {
        "MISSION_KEY_PERSONNEL",
        "EXPERIMENT_KEY_PERSONNEL",
        "INSTRUMENT_NAME",
        "INSTRUMENT_DESCRIPTION",
        "DATASET_CAVEATS",
        "METADATA_TYPE",
        "LOGICAL_FILE_ID",
        "EXPERIMENT_DESCRIPTION",
        "DATASET_ID",
        "MIN_TIME_RESOLUTION",
        "PROCESSING_LEVEL",
        "DATASET_VERSION",
        "FILE_TYPE",
        "FILE_CAVEATS",
        "TIME_RESOLUTION",
        "DATASET_DESCRIPTION",
        "MISSION_REFERENCES",
        "MISSION_CAVEATS",
        "EXPERIMENT",
        "OBSERVATORY_DESCRIPTION",
        "MEASUREMENT_TYPE",
        "ACKNOWLEDGEMENT",
        "METADATA_VERSION",
        "FILE_TIME_SPAN",
        "OBSERVATORY_CAVEATS",
        "OBSERVATORY_REGION",
        "INSTRUMENT_CAVEATS",
        "DATASET_TITLE",
        "INVESTIGATOR_COORDINATES",
        "EXPERIMENT_CAVEATS",
        "MISSION_TIME_SPAN",
        "MISSION_AGENCY",
        "MISSION_REGION",
        "EXPERIMENT_REFERENCES",
        "INSTRUMENT_TYPE",
        "DATA_TYPE",
        "MAX_TIME_RESOLUTION",
        "MISSION",
        "GENERATION_DATE",
        "MISSION_DESCRIPTION",
        "OBSERVATORY",
        "OBSERVATORY_TIME_SPAN",
        "CONTACT_COORDINATES",
        "VERSION_NUMBER",
    }
    
    for _, k := range mks {
        if entry, value_type, err := r_header.getMeta(k); err == nil {
            fmt.Println("Meta:ENTRY", entry)
            fmt.Println("Meta:VALUE_TYPE", value_type)
        } else {
            fmt.Println(err)
        }
        
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    //x r_header.dumpMeta()
    
    
	println("Lines read -> ", ix)

	//x r_header.m_data.dump()
    
    println("//x (cef_header.go) r_header.dump()")
	//x r_header.dump()

	return
}
