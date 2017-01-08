package main

import (
    "fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

func main() {

    fmt.Println(diag.BoldMagenta("ceflogs v0.0.1 (7 Jan 2016)"))
    fmt.Println(diag.BoldMagenta("Off we go!"))

	_, err := utils.NewCefArgs()

	if err != nil {
		fmt.Println(diag.BoldRed("Invalid Command Line Args"))	
		return
	}
    
//x	fmt.Println(diag.BoldRed(fmt.Sprintf("Error parsing cef file \n%#v", err.Error())))

    fmt.Println(diag.BoldMagenta("CAA Rocks!"))
}

