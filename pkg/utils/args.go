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
    m_trace bool
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



func (a1s *CefArgs) init() error {
	err := error(nil)

	flag.Var(&a1s.m_includes, "i", "Include Folders")
    flag.StringVar(&a1s.m_cefpath , "f", "", "Cef file path (.cef/.cef.gz)")
    flag.BoolVar(&a1s.m_trace , "t", false, "Show Trace debug  (false)")

	flag.Parse()

	diag.SetTrace(a1s.m_trace)

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		err = errors.New("Error: Invalid number of cmdline args")
	}

    s1 := a1s.m_cefpath
    ix := strings.LastIndexAny(s1, `\\/`); 
    if ix >= 0 {
        l:=0
        if strings.HasSuffix(s1, ".gz") {
            l = 3
        }
        
        a1s.m_filename = s1[ix+1:len(s1)-l]
    } else {
        a1s.m_filename = s1
    }
    
	a1s.Dump()

	return err
}

func (a1s CefArgs) Dump() {

//x     diag.Info(diag.BoldYellow("cef cefpath"), a1s.m_cefpath)
//x     diag.Info(diag.BoldYellow("cef filename"), a1s.m_filename)
//x     diag.Info(diag.BoldYellow("cef includes"), a1s.m_includes)
//x     diag.Info(diag.BoldYellow("trace"), a1s.m_trace)

    diag.Trace(diag.BoldYellow("cef cefpath"), a1s.m_cefpath)
    diag.Trace(diag.BoldYellow("cef filename"), a1s.m_filename)
    diag.Trace(diag.BoldYellow("cef includes"), a1s.m_includes)
    diag.Trace(diag.BoldYellow("trace"), a1s.m_trace)

}

func NewCefArgs() (CefArgs, error) {
	s_args = CefArgs{}
	err := s_args.init()
    
	return s_args, err
}

func GetCefArgs() (CefArgs) {
	return s_args		
}


