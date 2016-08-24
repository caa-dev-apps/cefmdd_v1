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

//x func enum_parser(v string, enums []string) (err error){ 
//x 
//x //x     for _, en := range kw.kw_enums {
//x     for _, e := range enums {
//x         if e == v {
//x            return
//x         }
//x     }
//x     
//x     return errors.New("Parse ENUM: Error not found")     
//x }

func formatted_parser(v string) (err error) { 
    fmt.Println("------- formatted_parser ---------> " + v)
    // "---->---->------" or "---->----"
    
    fs := strings.Split(v, ">")
    if len(fs) >= 2 {
        return
    }
    
    return errors.New("Parse FORMAT: Error")     
}

func integer_parser(v string) (err error) { 
    fmt.Println("------- integer_parser ---------> " + v)
    
	_, err = strconv.ParseInt(v, 10, 64)
    
    return
}

func iso_time_parser(v string) (err error) { 
    fmt.Println("------- iso_time_parser ---------> " + v)
    // "2012-04-11T15:57:15.012345678Z"
    _, err = time.Parse(time.RFC3339Nano, v);

    return 
}

func iso_time_range_parser(v string) (err error) { 
    fmt.Println("------- iso_time_range_parser ---------> " + v)
    // "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z"
    
    ts := strings.Split(v, "/")
    if len(ts) == 2 {
        if err = iso_time_parser(ts[0]); err != nil {
            return
        }
        
        return iso_time_parser(ts[1]);
    }
    
    return errors.New("Parse ISO_TIME_RANGE: Error")     
}

func numerical_parser(v string) (err error) { 
    fmt.Println("------- numerical_parser ---------> " + v)

	_, err = strconv.ParseFloat(v, 64)
    
    return
}

func string_parser(v string) (err error) { 
    fmt.Println("------- string_parser ---------> " + v)
    
    if len(v) > 0 {
        return 
    }
    
    return errors.New("Parse STRING: Empty") 
}

func text_parser(v string) (err error) { 
    fmt.Println("------- text_parser ---------> " + v)

    if len(v) > 0 {
        return 
    }
    
    return errors.New("Parse TEXT: Empty") 
}






