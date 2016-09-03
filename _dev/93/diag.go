package main

import (
	"fmt"
	"os"
	"text/tabwriter"
    //x "io"
)

//x func main() {
//x 	// Observe that the third line has no trailing tab,
//x 	// so its final cell is not part of an aligned column.
//x 	const padding = 3
//x 	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
//x 	fmt.Fprintln(w, "a\tb\taligned\t")
//x 	fmt.Fprintln(w, "aa\tbb\taligned\t")
//x 	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
//x 	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
//x 	w.Flush()
//x 
//x }

type Diag struct {
    m_writer *tabwriter.Writer
}

func NewDiag() *Diag {
    const padding = 3
//x     return &Diag{ m_writer: tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug) }
    return &Diag{ m_writer: tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight) }
}

func (d *Diag) Writeln(ss ...string) {

    for _, s := range ss {
        fmt.Fprint(d.m_writer, s, "\t")
    }
    fmt.Fprintln(d.m_writer, "")
    
}

func (d *Diag) Writeln2(ss []string) {

    for _, s := range ss {
        fmt.Fprint(d.m_writer, s, "\t")
    }
    fmt.Fprintln(d.m_writer, "")
    
}



func (d *Diag) Flush() {
    d.m_writer.Flush()
}
    
