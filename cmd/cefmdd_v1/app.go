package main

import (
    "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    fmt.Println(diag.BoldMagenta("cefmdd_v1 v0.0.4, (2 Jan 2016)"))
    fmt.Println(diag.BoldMagenta("Are we there yet?"))

	s_args, err := utils.NewCefArgs()

	if err != nil {
		fmt.Println(diag.BoldRed("Invalid Command Line Args"))	
		return
	}
    
	err = ReadCef(&s_args)
	if err != nil {
		fmt.Println(diag.BoldRed(fmt.Sprintf("Error parsing cef file \n%#v", err.Error())))
		return
	}

    fmt.Println(diag.BoldMagenta("CAA Rocks!"))
}
