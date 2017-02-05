package diag

import (
	"fmt"
    "github.com/fatih/color"
)

var s_trace bool 
var s_show_error_line bool

func SetTrace(d bool) {
    s_trace = d
}

func SetShowErrorLine(d bool) {
    s_show_error_line = d
}

var (
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

func Println(v ...interface{}) {
    if s_trace == true {
        fmt.Println(v...)
    }
}

func Trace(tag string, v ...interface{}) {
    if s_trace == true {
        fmt.Print(tag, " ")
        fmt.Println(v...)
    }
}


func Info(tag string, v ...interface{}) {
    fmt.Print(Yellow(tag), " ")
    fmt.Println(v...)
}

func Todo(tag string, v ...interface{}) {
    fmt.Print(BoldMagenta("Todo: ", tag), " ")
    fmt.Println(v...)
}

func Warn(tag string, v ...interface{}) {
    fmt.Print(Magenta("Warn: ", tag), " ")
    fmt.Println(v...)
}

func Error(tag string, v ...interface{}) {
    fmt.Print(Red("Error: ", tag), " ")
    fmt.Println(v...)
}

func ErrorLine(v ...interface{}) {
    if s_show_error_line == true {
        fmt.Print("  ")
        fmt.Println(v...)
    }
}


func Fatal(tag string, v ...interface{}) {
    fmt.Print(BoldRed("Fatal: ", tag), " ")
    fmt.Println(v...)
}


func Infof(tag, format string, v ...interface{}) {
    Info(tag, fmt.Sprintf(format, v))
}

func Warnf(tag, format string, v ...interface{}) {
    Warn(tag, fmt.Sprintf(format, v))
}

func Errorf(tag, format string, v ...interface{}) {
    Error(tag, fmt.Sprintf(format, v))
}

func Fatalf(tag, format string, v ...interface{}) {
    Fatal(tag, fmt.Sprintf(format, v))
}
