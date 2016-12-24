package main

import (
    "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/diag"
    "github.com/caa-dev-apps/cefmdd_v1/utils"
)

//x var (
//x 	s_args CefArgs
//x )

func main() {

    fmt.Println(diag.BoldMagenta("cefmdd_v1 v0.0.2, (Dec 2016)"))

	s_args, err := utils.NewCefArgs()
	if err != nil {
		fmt.Println(diag.BoldRed("Invalid Command Line Args"))	
		return
	}
    
	err = ReadCef(&s_args)
	if err != nil {
		fmt.Println(diag.BoldRed("Error parsing cef file"))	
		return
	}

    fmt.Println(diag.BoldMagenta("CAA Rocks!"))


}
