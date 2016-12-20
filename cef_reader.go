package main

import (
    "fmt"
)


func ReadCef(args *CefArgs) (err error) {

	l_path := args.m_cefpath
	l_lines := EachLine(l_path)

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

