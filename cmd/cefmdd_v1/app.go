package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.BoldMagenta("cefmdd_v1 v0.0.5, (4 Jan 2016)"))
    diag.Info(diag.BoldMagenta("Are we there yet?"))

	s_args, err := utils.NewCefArgs()

	if err != nil {
		diag.Error(diag.BoldRed("Invalid Command Line Args"))	
		return
	}

	err = ReadCef(&s_args)
	if err != nil {
		//x diag.Error(diag.BoldRed(fmt.Sprintf("Error parsing cef file \n%#v", err.Error())))
		diag.Errorf(diag.BoldRed("Error parsing cef file"), "\n%#v", err.Error())
		return
	}

    diag.Info(diag.BoldMagenta("CAA Rocks!"))
}
