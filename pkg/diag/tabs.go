package diag

import (
	"fmt"
	"os"
	"text/tabwriter"
//x     "github.com/fatih/color"
)

type Tabs struct {
    m_writer *tabwriter.Writer
}


//x var (
//x     //x info2 := color.New(color.FgWhite, color.BgCyan).SprintFunc()
//x     BoldWhite   = color.New(color.FgWhite, color.Bold).SprintFunc()
//x     BoldMagenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
//x     BoldRed     = color.New(color.FgRed, color.Bold).SprintFunc()
//x     BoldBlue    = color.New(color.FgBlue, color.Bold).SprintFunc()
//x     BoldGreen   = color.New(color.FgGreen, color.Bold).SprintFunc()
//x     BoldYellow  = color.New(color.FgYellow, color.Bold).SprintFunc()
//x     BoldCyan    = color.New(color.FgCyan, color.Bold).SprintFunc()
//x )


func NewTabs() *Tabs {
    const padding = 3
    return &Tabs{ m_writer: tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight) }
}

func (d *Tabs) Printf(format string, a ...interface{}) {
    fmt.Fprintf(d.m_writer, format, a...)
}

func (d *Tabs) Println(ss ...string) {

    for _, s := range ss {
        fmt.Fprint(d.m_writer, s, "\t")
    }
    fmt.Fprintln(d.m_writer, "")
    
}

//x func (d *Tabs) PrintlnN(n int, ss ...string) {
//x 
//x     for i:=0; i<n; i++ {
//x         fmt.Fprint(d.m_writer, "", "\t")
//x     }
//x 
//x     d.Println(ss...)
//x }
//x 
//x func (d *Tabs) Flush() {
//x     d.m_writer.Flush()
//x }
    
