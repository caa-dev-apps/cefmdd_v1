package diag

import (
	"fmt"
    "github.com/fatih/color"
)

var s_debug bool 

func SetDebug(d bool) {
    s_debug = d
}


var (
    //x info2 := color.New(color.FgWhite, color.BgCyan).SprintFunc()
    BoldWhite   = color.New(color.FgWhite, color.Bold).SprintFunc()
    BoldMagenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
    BoldRed     = color.New(color.FgRed, color.Bold).SprintFunc()
    BoldBlue    = color.New(color.FgBlue, color.Bold).SprintFunc()
    BoldGreen   = color.New(color.FgGreen, color.Bold).SprintFunc()
    BoldYellow  = color.New(color.FgYellow, color.Bold).SprintFunc()
    BoldCyan    = color.New(color.FgCyan, color.Bold).SprintFunc()
)

func Trace(tag string, ss ...string) {
    if s_debug == true {
        fmt.Println(tag, ss)
    }
}

func Info(tag string, ss ...string) {
    fmt.Println(tag, ss)
}

func Warn(tag string, ss ...string) {
    fmt.Println(tag, ss)
}

func Error(tag string, ss ...string) {
    fmt.Println(tag, ss)
}

func Fatal(tag string, ss ...string) {
    fmt.Println(tag, ss)
}




