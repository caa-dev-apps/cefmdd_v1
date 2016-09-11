package main

import (
    "fmt"
)

func mooi_log(a ...interface{}) (n int, err error) {
	return
}

func main() {

    fmt.Println(BoldMagenta("cefmdd_v1 v0.0.1, (Sept 2016)"))
    //x fmt.Println(BoldWhite("User Home: " + UserHomeDir()))

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	_, err = ReadCefHeader(&args)
	error_check(err, "Error parsing header")
    
    fmt.Println(BoldMagenta("Go CAA-Team!"))
}
