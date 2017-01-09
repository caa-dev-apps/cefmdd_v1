package diag

import (
	"fmt"
    "github.com/fatih/color"
)

var s_info bool 

func SetInfo(d bool) {
    s_info = d
}


var (
    //x info2 := color.New(color.FgWhite, color.BgCyan).SprintFunc()

    White       = color.New(color.FgWhite).SprintFunc()
    Magenta     = color.New(color.FgMagenta).SprintFunc()
    Red         = color.New(color.FgRed).SprintFunc()
    Blue        = color.New(color.FgBlue).SprintFunc()
    Green       = color.New(color.FgGreen).SprintFunc()
    Yellow      = color.New(color.FgYellow).SprintFunc()
    Cyan        = color.New(color.FgCyan).SprintFunc()

    BoldWhite   = color.New(color.FgWhite, color.Bold).SprintFunc()
    BoldMagenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
    BoldRed     = color.New(color.FgRed, color.Bold).SprintFunc()
    BoldBlue    = color.New(color.FgBlue, color.Bold).SprintFunc()
    BoldGreen   = color.New(color.FgGreen, color.Bold).SprintFunc()
    BoldYellow  = color.New(color.FgYellow, color.Bold).SprintFunc()
    BoldCyan    = color.New(color.FgCyan, color.Bold).SprintFunc()

)

//x func Info(tag string, ss ...string) {
//x     if s_info == true {
//x         fmt.Println(tag, ss)
//x     }
//x }
//x 
//x func Warn(tag string, ss ...string) {
//x     fmt.Println(tag, ss)
//x }
//x 
//x func Error(tag string, ss ...string) {
//x     fmt.Println(tag, ss)
//x }
//x 
//x func Fatal(tag string, ss ...string) {
//x     fmt.Println(tag, ss)
//x }


func Info(tag string, v ...interface{}) {
    if s_info == true {
        fmt.Println(tag, v)
    }
}

func Warn(tag string, v ...interface{}) {
    fmt.Println(tag, v)
}

func Error(tag string, v ...interface{}) {
    fmt.Println(tag, v)
}

func Fatal(tag string, v ...interface{}) {
    fmt.Println(tag, v)
}



func Infof(tag, format string, v ...interface{}) {
    if s_info == true {
        fmt.Println(tag, fmt.Sprintf(format, v))
    }
}

func Warnf(tag, format string, v ...interface{}) {
    fmt.Println(tag, fmt.Sprintf(format, v))
}

func Errorf(tag, format string, v ...interface{}) {
    fmt.Println(tag, fmt.Sprintf(format, v))
}

func Fatalf(tag, format string, v ...interface{}) {
    fmt.Println(tag, fmt.Sprintf(format, v))
}

