package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.Yellow("cefmdd v1.0.1, (17 Jan 2017)"))

	s_args, err := utils.NewCefArgs()

	if err != nil {
		//x diag.Error(diag.BoldRed("Invalid Command Line Args"))	
		return
	}

	err = ReadCef(&s_args)
	if err != nil {
		diag.Errorf(diag.BoldRed("Error parsing cef file"), "\n%#v", err.Error())
		return
	}

    diag.Info(diag.Yellow("CAA Rocks!\n"))
}
