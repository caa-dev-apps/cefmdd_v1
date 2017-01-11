package diag

import (
	"fmt"
    "github.com/fatih/color"
)

var s_trace bool 

func SetTrace(d bool) {
    s_trace = d
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


func Trace(tag string, v ...interface{}) {
    if s_trace == true {
        fmt.Println(tag, v)
    }
}

func Info(tag string, v ...interface{}) {
        fmt.Println(tag, v)
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



func Tracef(tag, format string, v ...interface{}) {
    if s_trace == true {
        fmt.Println(tag, fmt.Sprintf(format, v))
    }
}

func Infof(tag, format string, v ...interface{}) {
    fmt.Println(tag, fmt.Sprintf(format, v))
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

