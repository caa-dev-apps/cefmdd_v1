package utils

import (
    "fmt"
    "os"
    "time"
    "strconv"
    "errors"
    "strings"
    "runtime"
)

///////////////////////////////////////////////////////////////////////////////
//

func UserHomeDir() string {
    
    //x pwd, _ := os.Getwd()
    //x fmt.Println("Getwd : ", pwd)
    
    if runtime.GOOS == "windows" {
        home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
        
        if home == "" {
            home = os.Getenv("USERPROFILE")
        }
        return home
    }
    return os.Getenv("HOME")
}

func CefMddDir() string {
//x    return UserHomeDir() + `/.cefmdd_v1`
    return UserHomeDir() + `/_cefmdd_v1`
}


func Error_check(err error, i_s string) {
    if err != nil {
        fmt.Println(err.Error())
        fmt.Println("Error-Check: ", i_s)
        os.Exit(1)
    }
}

func FileExists(name string) (isReq bool, err error) {

    info, err := os.Stat(name)

    if err == nil {
        isReq = info.Mode().IsRegular()
    }

    return
}

///////////////////////////////////////////////////////////////////////////////
//

func Is_quoted_string(s string) (r bool) {

    l := len(s)
    r = l >= 2 && s[0] == '"' && s[l-1] == '"'

    return
}

func Trim_quoted_string(s string) (string) {
    return strings.Trim(s, `"`)
}

///////////////////////////////////////////////////////////////////////////////
//

func On_enum_not_found_error(i_keyword string) (err error) {
    return errors.New("Parse Error: Enum keyword not found, " + i_keyword)     
}


func On_enum_parser_error(i_keyword, i_value string) (err error) {
    return errors.New("Parse Error: Enum value not found, " + i_keyword + " -> " + i_value)     
}

func On_parser_error(i_parser, i_value string) (err error) {
    return errors.New("Parse Error: " + i_parser + " -> " + i_value)     
}

///////////////////////////////////////////////////////////////////////////////
//

func Iso_time(v string) (time.Time, error) { 
    // "2012-04-11T15:57:15.012345678Z"
    return time.Parse(time.RFC3339Nano, v)
}

///////////////////////////////////////////////////////////////////////////////
//

func Formatted_parser(v string) (err error) { 
    
    fs := strings.Split(v, ">")
    if len(fs) < 2 {
        return On_parser_error("Formatted", v)
    }
    
    return 
}

func Integer_parser(v string) (err error) { 
    
    if _, err = strconv.ParseInt(v, 10, 64); err != nil {
        return On_parser_error("Integer", v)
    }
    
    return
}

func Iso_time_parser(v string) (err error) { 
    // "2012-04-11T15:57:15.012345678Z"
    if _, err = time.Parse(time.RFC3339Nano, v); err != nil {
        return On_parser_error("ISO Time", v)
    }

    return 
}

func Iso_time_range_parser(v string) (err error) { 
    // "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z"
    
    ts := strings.Split(v, "/")
    if len(ts) == 2 {
        if err = Iso_time_parser(ts[0]); err == nil {
            if err = Iso_time_parser(ts[1]); err == nil {
                return 
            }
        }
    }
    
    return On_parser_error("ISO Time Range", v)
}

func Numerical_parser(v string) (err error) { 

    if _, err = strconv.ParseFloat(v, 64); err != nil {
        return On_parser_error("Numerical", v)
    }
    
    return
}

func String_parser(v string) (err error) { 
    
//x     if len(v) == 0 {
//x         return On_parser_error("String", v)
//x     }
    
    return 
}

func Text_parser(v string) (err error) { 

//x     if len(v) == 0 {
//x         return On_parser_error("String", v)
//x     }
    
    return 
}

func String_matcher_test(v0 string) (f func(v1 string) (err error)) { 

    return func(v string) (err error) {
        if v0 != v {
            err = errors.New(fmt.Sprintf("String Mismatch: Expected %s : Received %s", v0, v))
        }

        return
    }

    return 
}
