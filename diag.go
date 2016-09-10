package main

import (
	"fmt"
	"os"
	"text/tabwriter"
    "github.com/fatih/color"

)

type Diag struct {
    m_writer *tabwriter.Writer
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


func NewDiag() *Diag {
    const padding = 3
    //- return &Diag{ m_writer: tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug) }  
    return &Diag{ m_writer: tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight) }
}

func (d *Diag) Printf(format string, a ...interface{}) {
    fmt.Fprintf(d.m_writer, format, a...)
}

func (d *Diag) Println(ss ...string) {

    for _, s := range ss {
        fmt.Fprint(d.m_writer, s, "\t")
    }
    fmt.Fprintln(d.m_writer, "")
    
}

func (d *Diag) PrintlnN(n int, ss ...string) {

    for i:=0; i<n; i++ {
        fmt.Fprint(d.m_writer, "", "\t")
    }

    d.Println(ss...)
}

func (d *Diag) Flush() {
    d.m_writer.Flush()
}
    
