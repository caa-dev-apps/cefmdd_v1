package main

import (
	"errors"
	"flag"
	"fmt"
//x     "path"
    "strings"
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
	m_cefpath  *string
    m_filename string
}

//x *s_args.m_cefpath
func (a1s *CefArgs) init() error {
	err := error(nil)

	flag.Var(&a1s.m_includes, "i", "Include Folders")
	a1s.m_cefpath = flag.String("f", "", "Cef file path (.cef/.cef.gz)")
    
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		err = errors.New("Error: Invalid number of cmdline args")
	}

    s1 := *a1s.m_cefpath
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
    
	return err
}

func (a1s *CefArgs) dump() {
	mooi_log("cef path: ", *a1s.m_cefpath)
	mooi_log("include folders", a1s.m_includes)
}

func NewCefArgs() (args CefArgs, err error) {
	args = CefArgs{}
	err = args.init()
    
	return args, err
}
