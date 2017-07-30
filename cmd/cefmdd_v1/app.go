package main

import (
	"os"

    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.Yellow("cefmdd v1.1.1, (29 July 2017)"))

	s_args, err := utils.NewCefArgs()
	if err != nil {
		os.Exit(-1)
	}

	err = ReadCef(&s_args)
	if err != nil {
		diag.Errorf(diag.BoldRed("Error parsing cef file"), "\n%v", err.Error())
		os.Exit(-1)
	}

    diag.Info(diag.Yellow("CAA Rocks!\n"))

    return
}
