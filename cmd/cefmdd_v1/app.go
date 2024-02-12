package main

import (

	"os"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

//20200120 added check for matched quotes
//20200123 added check that CC files match CONTEXT OBS
//20200123 added check that attributes must be unique within param
//20200207 added check for EOR marker on last data line
//20200207 updated observatory from context to cluster-context
//20210615    diag.Info(diag.Yellow("cefmdd v1.2.5 BETA, (20200702)"))
//20230929    diag.Info(diag.Yellow("cefmdd v1.2.7 BETA, (20230301)"))
//2024 test
        diag.Info(diag.Yellow("cefmdd v1.2.7b, (20230929)"))
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
