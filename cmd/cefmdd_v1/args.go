package main

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
    
	a1s.dump()

	return err
}

func (a1s *CefArgs) dump() {

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


func NewCefArgs() (s_args CefArgs, err error) {
	s_args = CefArgs{}
	err = s_args.init()
    
	return s_args, err
}









