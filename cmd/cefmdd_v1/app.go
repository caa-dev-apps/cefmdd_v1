package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.Yellow("\ncefmdd_v1 v0.0.7, (14 Jan 2017)"))

	s_args, err := utils.NewCefArgs()

	if err != nil {
		diag.Error(diag.BoldRed("Invalid Command Line Args"))	
		return
	}

	err = ReadCef(&s_args)
	if err != nil {
		diag.Errorf(diag.BoldRed("Error parsing cef file"), "\n%#v", err.Error())
		return
	}

    diag.Info(diag.Yellow("CAA Rocks!"))
}
