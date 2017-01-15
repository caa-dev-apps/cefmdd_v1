package main

import (
    //x "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    args, err := utils.NewCefArgs()

    diag.Trace("TEST!",     "__Trace__")
    diag.Info("TEST!",      "__Info___")
    diag.Warn("TEST!",      "__Warn___")
    diag.Error("TEST!",     "__Error__")
    diag.Fatal("TEST!",     "__Fatal__")


    diag.Trace(diag.BoldMagenta("ceflogs v0.0.1 (7 Jan 2017)"))
    diag.Trace(diag.BoldMagenta("Off we go!"))
    args.Dump()

	if err != nil {
		diag.Error(diag.BoldRed("Invalid Command Line Args"))	
		return
	}

    diag.Trace(diag.BoldMagenta("CAA Rocks!"))
}

