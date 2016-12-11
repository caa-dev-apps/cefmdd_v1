package main

import (
    "fmt"
//x     "errors"
)

func mooi_log(a ...interface{}) (n int, err error) {
	return
}

var (
	s_args CefArgs
)

//- func main() {
//- 
//-     fmt.Println(BoldMagenta("cefmdd_v1 v0.0.1, (Sept 2016)"))
//- 
//- 	args, err := NewCefArgs()
//- 	error_check(err, "Invalid Command Line Args")
//- 	args.dump()
//- 
//- 	_, err = ReadCefHeader(&args)
//- 	error_check(err, "Error parsing header")
//-     
//-     fmt.Println(BoldMagenta("CAA Rocks!"))
//- }

func main() {

    fmt.Println(BoldMagenta("cefmdd_v1 v0.0.2, (Dec 2016)"))

    var err error
	s_args, err = NewCefArgs()
    
    fmt.Println(BoldYellow("cef cefpath"), *s_args.m_cefpath)
    fmt.Println(BoldYellow("cef filename"), s_args.m_filename)
    
	error_check(err, "Invalid Command Line Args")
	s_args.dump()

	//x _, err = ReadCefHeader(&s_args)
	//x error_check(err, "Error parsing header")
    
	err = ReadCef(&s_args)
	error_check(err, "Error parsing cef file")


    fmt.Println(BoldMagenta("CAA Rocks!"))
}
