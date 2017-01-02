package utils

import (
	"errors"
	"flag"
	"fmt"
    "strings"
    "github.com/caa-dev-apps/cefmdd_v1/diag"
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


func (a1s *CefArgs) init() error {
	err := error(nil)

	flag.Var(&a1s.m_includes, "i", "Include Folders")
    flag.StringVar(&a1s.m_cefpath , "f", "", "Cef file path (.cef/.cef.gz)")
	flag.Parse()

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

    fmt.Println(diag.BoldYellow("cef cefpath"), a1s.m_cefpath)
    fmt.Println(diag.BoldYellow("cef filename"), a1s.m_filename)	
    fmt.Println(diag.BoldYellow("cef includes"), a1s.m_includes)	
}

//x func NewCefArgs() (args CefArgs, err error) {
//x 	args = CefArgs{}
//x 	err = args.init()
//x     
//x 	return args, err
//x }
//x 
//x s_args


//d func NewCefArgs() (r_args CefArgs, err error) {
//d 	s_args = CefArgs{}
//d 	err = s_args.init()
//d     
//d 	r_args = s_args
//d 
//d 	return 
//d }

func NewCefArgs() (CefArgs, error) {
	s_args = CefArgs{}
	err := s_args.init()
    
	return s_args, err
}

func GetCefArgs() (CefArgs) {
	return s_args		
}


