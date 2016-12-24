package main

import (
    "fmt"
	"github.com/caa-dev-apps/cefmdd_v1/readers"    
	"github.com/caa-dev-apps/cefmdd_v1/utils"    
)


func ReadCef(args *utils.CefArgs) (err error) {

//x 	l_path := args.m_cefpath
	l_path := args.GetCefPath()
	l_lines := readers.EachLine(l_path)

	fmt.Println("....", l_path)

//x 	fmt.Println("l_path: ", l_path)
//x 	return
//x 	for l := range l_lines {
//x 
//x 	}

	_, err = ReadHeader(args, l_lines)
	if err != nil {
		return
	}

	err = ReadData(args, l_lines)

	return
}

