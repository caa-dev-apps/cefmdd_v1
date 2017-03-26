package utils

import (
    "fmt"
    "os"
    "time"
    "strconv"
    "errors"
    "strings"
    "runtime"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

///////////////////////////////////////////////////////////////////////////////
//

func UserHomeDir() string {
    
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
    return UserHomeDir() + `/_cefmdd_v1`
}


func Error_check(err error, i_s string) {
    if err != nil {
        diag.Error(err.Error())
        diag.Error("Error-Check: ", i_s)
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

func SizesProduct(sizes []string) (product int64, err error) {

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
    return errors.New("Parser, Enum keyword not found")     
}

func On_enum_parser_error(i_keyword, i_value string) (err error) {
    return errors.New("Parser, Enum value not found ")     
}

func On_enum_description_error(i_keyword, i_value string) (err error) {
    return errors.New("Parser, Enum description mismatch ")     
}


func On_parser_error(i_parser, i_value string) (err error) {
    return errors.New("Parser, " + i_parser + " -> " + i_value)     
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

func Get_time_range(v string) (t0, t1 time.Time, err error) { 



    ts := strings.Split(v, "/")
    if len(ts) == 2 {
        t0, err = time.Parse(time.RFC3339Nano, ts[0])
        if err == nil {
            t1, err = time.Parse(time.RFC3339Nano, ts[1])
            if err == nil {
                if t0.Before(t1) {
                    return
                }
            }
        }
    }

    err = On_parser_error("ISO Time Range", v)

    return 
}


func Iso_time_range_parser(v string) (err error) { 
    // "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z"

    t0, t1, err := Get_time_range(v)
    
    if err == nil {
        if t0.Before(t1) {
            return
        }        
    }

    return On_parser_error("ISO Time Range", v)
}

func Iso_time_range_parser_allow_equal(v string) (err error) { 
    // "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z"

    t0, t1, err := Get_time_range(v)
    
    if err == nil {
        if t0.Before(t1) || t0.Equal(t1) {
            return
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
    
    return 
}

func Text_parser(v string) (err error) { 

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

///////////////////////////////////////////////////////////////////////////////
//

func Is_Numerical(v string) (bool) { 

    return Numerical_parser(v) == nil
}

///////////////////////////////////////////////////////////////////////////////
//

func ValueTypeParserFunc(vt string) (f func(string) (error), err error) {

    switch vt {
        case "CHAR" :               f = String_parser
        case "INT" :                f = Integer_parser
        case "FLOAT" :              f = Numerical_parser
        case "DOUBLE" :             f = Numerical_parser
        case "ISO_TIME" :           f = Iso_time_parser
        case "ISO_TIME_RANGE" :     f = Iso_time_range_parser_allow_equal
        default :
            err = errors.New("Unknown VALUE_TYPE : " + vt)
    }

    return
}


///////////////////////////////////////////////////////////////////////////////
//

// e.g. filename = CN_CP_EDI_SPIN__2001....
// return observatory, data_type, experiment, error
// i.e CN, CP, EDI err

func GetFirst3CefFilenameParts() (observatory, data_type, experiment string, err error) {

    l_filename := GetCefArgs().GetFilename()

    if len(l_filename) > 7 && l_filename[2] == '_' && l_filename[5] == '_' {
        observatory = l_filename[0:2]
        data_type = l_filename[3:5]
        ix := strings.Index(l_filename[6:], "_")
        if ix > 0 {
            experiment = l_filename[6:6+ix]
            return
        }
    } 

    diag.Error("Maformed filename: ", l_filename)
    err = errors.New("Maformed filename: " + l_filename)

    return 
}
