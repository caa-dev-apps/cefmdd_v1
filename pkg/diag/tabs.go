package diag

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Tabs struct {
    m_writer *tabwriter.Writer
}

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
