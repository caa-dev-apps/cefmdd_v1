package main

import (
//	"fmt"
	"github.com/caa-dev-apps/cefmdd_v1/pkg/readers"    
	"github.com/caa-dev-apps/cefmdd_v1/pkg/utils"    
)


func ReadCef(args *utils.CefArgs) (err error) {

	l_path := args.GetCefPath()
	l_lines := readers.EachLine(l_path)
	l_header, err := ReadHeader(args, l_lines)
	if err != nil {
		return err
	}

	err = ReadData(l_header, l_lines)

	return
}

