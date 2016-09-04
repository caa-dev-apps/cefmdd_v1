package main

import (
	"fmt"
	"os"
	"text/tabwriter"
    //x "io"
)

type Diag struct {
    m_writer *tabwriter.Writer
}

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
    
