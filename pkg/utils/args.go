package utils

import (
	"errors"
	"flag"
	"fmt"
    "strings"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

var (
	s_args CefArgs
)

type strslice []string

func (ss *strslice) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *strslice) Set(s string) error {
	*ss = append(*ss, s)
	return nil
}

type CefArgs struct {
	m_includes strslice
	m_cefpath  string
    m_filename string
    m_max_records int
    m_trace bool
    m_warn bool
    m_duplicates bool
    m_ignore_quotes bool
    m_show_error_line bool
}

func (a1s CefArgs) GetIncludes() (strslice) {
	return a1s.m_includes	
}

func (a1s CefArgs) GetCefPath() (string) {
	return a1s.m_cefpath
}

func (a1s CefArgs) GetFilename() (string) {
	return a1s.m_filename
}

func (a1s CefArgs) GetInfo() (bool) {
	return a1s.m_trace
}

func (a1s CefArgs) GetDuplicateFlag() (bool) {
	return a1s.m_duplicates
}

func (a1s CefArgs) GetIgnoreQuotesFlag() (bool) {
	return a1s.m_ignore_quotes
}

func (a1s CefArgs) GetMaxRecords() (int) {
	return a1s.m_max_records
}



func (a1s *CefArgs) init() (err error) {

	flag.Var(&a1s.m_includes, "i", "Include Folders")
    flag.StringVar(&a1s.m_cefpath , "f", "", "Cef file path (.cef/.cef.gz)")
    flag.BoolVar(&a1s.m_trace , "t", false, "Show trace debug")
    flag.BoolVar(&a1s.m_warn , "w", false, "Show warnings")
    flag.BoolVar(&a1s.m_duplicates , "d", false, "Allow duplicate timestamps")
 //   flag.BoolVar(&a1s.m_ignore_quotes , "q", false, "Used for regression testing")
     flag.BoolVar(&a1s.m_ignore_quotes , "q", false, "")
    flag.BoolVar(&a1s.m_show_error_line , "e", true, "Show error line details")
//  flag.IntVar(&a1s.m_max_records , "n", 100, "Max number of records to read (all -1)")

    flag.IntVar(&a1s.m_max_records , "n", -1, "Max number of records to read (all -1)")

	flag.Parse()

	diag.SetTrace(a1s.m_trace)
	diag.SetWarn(a1s.m_warn)
	diag.SetShowErrorLine(a1s.m_show_error_line)
	diag.SetDuplicates(a1s.m_duplicates)
	diag.SetIgnoreQuotes(a1s.m_ignore_quotes)

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		err = errors.New("Error: Invalid number of cmdline args")
	}

    s1 := a1s.m_cefpath
    ix := strings.LastIndexAny(s1, `\\/`); 
    if ix >= 0 {
        a1s.m_filename = s1[ix+1:len(s1)]
    } else {
        a1s.m_filename = s1
    }

    if strings.HasSuffix(a1s.m_filename, ".gz") {
        a1s.m_filename = a1s.m_filename[0:len(a1s.m_filename)-3]
    }
    
	a1s.Dump()

	return err
}

func (a1s CefArgs) Dump() {

    diag.Trace(diag.BoldYellow("cef cefpath"), a1s.m_cefpath)
    diag.Trace(diag.BoldYellow("cef filename"), a1s.m_filename)
    diag.Trace(diag.BoldYellow("cef includes"), a1s.m_includes)
    diag.Trace(diag.BoldYellow("m_max_records"), a1s.m_max_records)
    diag.Trace(diag.BoldYellow("trace"), a1s.m_trace)
    diag.Trace(diag.BoldYellow("warn"), a1s.m_warn)
    diag.Trace(diag.BoldYellow("allow duplicates"), a1s.m_duplicates)
    diag.Trace(diag.BoldYellow("ignore quotes"), a1s.m_ignore_quotes)
    diag.Trace(diag.BoldYellow("show error line"), a1s.m_show_error_line)

}

func NewCefArgs() (CefArgs, error) {
	s_args = CefArgs{}
	err := s_args.init()
    
	return s_args, err
}

func GetCefArgs() (CefArgs) {
	return s_args		
}
