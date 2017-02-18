package main

import (
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    diag.Info(diag.Yellow("cefmdd v1.0.15, (18 Feb 2017)"))

	s_args, err := utils.NewCefArgs()
	if err != nil {
		return
	}

	err = ReadCef(&s_args)
	if err != nil {
		diag.Errorf(diag.BoldRed("Error parsing cef file"), "\n%v", err.Error())
		return
	}

    diag.Info(diag.Yellow("CAA Rocks!\n"))
}
