package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    args, err := utils.NewCefArgs()

    diag.Info(diag.BoldMagenta("ceflogs v0.0.1 (7 Jan 2016)"))
    diag.Info(diag.BoldMagenta("Off we go!"))
    args.Dump()


    diag.Trace(diag.BoldBlue("Trace"), "1", "2", "3")
    diag.Info(diag.BoldGreen("Info"), "1", "2", "3")
    diag.Warn(diag.BoldBlue("Warn"), "1", "2", "3")
    diag.Error(diag.BoldRed("Error"), "1", "2", "3")
    diag.Fatal(diag.BoldMagenta("Fatal"), "1", "2", "3")


	if err != nil {
		diag.Error(diag.BoldRed("Invalid Command Line Args"))	
		return
	}

    diag.Trace(diag.BoldMagenta("CAA Rocks!"))
}

