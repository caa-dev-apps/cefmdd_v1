package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.Yellow("cefmdd_v1 v0.0.6, (14 Jan 2016)"))

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

//x     diag.Trace("TEST!", 	"__Trace__")
//x     diag.Info("TEST!", 		"__Info___")
//x     diag.Warn("TEST!", 		"__Warn___")
//x     diag.Error("TEST!", 	"__Error__")
//x     diag.Fatal("TEST!", 	"__Fatal__")

    diag.Info(diag.Yellow("CAA Rocks!"))
}
