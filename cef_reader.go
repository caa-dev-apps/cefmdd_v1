package main

import (
    "fmt"
	"github.com/caa-dev-apps/cefmdd_v1/readers"    
	"github.com/caa-dev-apps/cefmdd_v1/utils"    
)


func ReadCef(args *utils.CefArgs) (err error) {

	l_path := args.GetCefPath()
	l_lines := readers.EachLine(l_path)

	fmt.Println("....", l_path)

	l_header, err := ReadHeader(args, l_lines)
	if err != nil {
		return
	}

	//x fmt.Printf("%#v", l_header)
	err = ReadData(l_header, l_lines)

	//x l_header.Dump()

	return
}

