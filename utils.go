package main

import (
	"fmt"
	"os"
    "time"
    "strconv"
    "errors"
    "strings"
)

func error_check(err error, i_s string) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error-Check: ", i_s)
		os.Exit(1)
	}
}

func fileExists(name string) (isReq bool, err error) {

	info, err := os.Stat(name)

	if err == nil {
		isReq = info.Mode().IsRegular()
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
//


func on_enum_not_found_error(i_keyword string) (err error) {
    return errors.New("Parse Error: Enum keyword not found, " + i_keyword)     
}


func on_enum_parser_error(i_keyword, i_value string) (err error) {
    return errors.New("Parse Error: Enum value not found, " + i_keyword + " -> " + i_value)     
}



func on_parser_error(i_parser, i_value string) (err error) {
    return errors.New("Parse Error: " + i_parser + " -> " + i_value)     
}

func formatted_parser(v string) (err error) { 
    
    fs := strings.Split(v, ">")
    if len(fs) < 2 {
        return on_parser_error("Formatted", v)
    }
    
    return 
}

func integer_parser(v string) (err error) { 
    
	if _, err = strconv.ParseInt(v, 10, 64); err != nil {
        return on_parser_error("Integer", v)
    }
    
    return
}

func iso_time_parser(v string) (err error) { 
    // "2012-04-11T15:57:15.012345678Z"
    if _, err = time.Parse(time.RFC3339Nano, v); err != nil {
        return on_parser_error("ISO Time", v)
    }

    return 
}

func iso_time_range_parser(v string) (err error) { 
    // "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z"
    
    ts := strings.Split(v, "/")
    if len(ts) == 2 {
        if err = iso_time_parser(ts[0]); err == nil {
            if err = iso_time_parser(ts[1]); err == nil {
                return 
            }
        }
    }
    
    return on_parser_error("ISO Time Range", v)
}

func numerical_parser(v string) (err error) { 

	if _, err = strconv.ParseFloat(v, 64); err != nil {
        return on_parser_error("Numerical", v)
    }
    
    return
}

func string_parser(v string) (err error) { 
    
//x     if len(v) == 0 {
//x         return on_parser_error("String", v)
//x     }
    
    return 
}

func text_parser(v string) (err error) { 

//x     if len(v) == 0 {
//x         return on_parser_error("String", v)
//x     }
    
    return 
}

