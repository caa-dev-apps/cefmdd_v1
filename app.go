package main

import (
    "fmt"
)

var (
	s_args CefArgs
)

func main() {

    fmt.Println(BoldMagenta("cefmdd_v1 v0.0.2, (Dec 2016)"))

	s_args, err := NewCefArgs()
	if err != nil {
		fmt.Println(BoldRed("Invalid Command Line Args"))	
		return
	}
    
	err = ReadCef(&s_args)
	if err != nil {
		fmt.Println(BoldRed("Error parsing cef file"))	
		return
	}

    fmt.Println(BoldMagenta("CAA Rocks!"))
}
