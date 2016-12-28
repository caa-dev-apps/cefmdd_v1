package main

import (
    "fmt"
	//x "github.com/caa-dev-apps/cefmdd_v1/header"    
	"github.com/caa-dev-apps/cefmdd_v1/readers"    
	"github.com/caa-dev-apps/cefmdd_v1/utils"    
)


func ReadCef(args *utils.CefArgs) (err error) {

	l_path := args.GetCefPath()
	l_lines := readers.EachLine(l_path)

	fmt.Println("....", l_path)

	//x _, err = ReadHeader(args, l_lines)
	l_header, err := ReadHeader(args, l_lines)
	if err != nil {
		return
	}

	//x fmt.Printf("%#v", l_header)
	err = ReadData(args, l_lines)

	//x l_header.Dump()

	vs := l_header.Vars().List()

//x 	for ix, v := range vs {
//x 		fmt.Println(ix)		
//x 		fmt.Println(v)		
//x 		fmt.Println("\n\n")		
//x 	}

	//x vs,			 []Attrs
	//x Attrs v, 	map[string][]string
	for ix, v := range vs {
		fmt.Println(ix)		

		for k1, v1 := range v.Map() {
			fmt.Println(k1, "\t", v1)		
		}

		//x fmt.Println(v)		
		fmt.Println("\n\n")		
	}

	//x fmt.Println(vs)

	return
}

