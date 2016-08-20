package main

import (
)

func mooi_log(a ...interface{}) (n int, err error) {
	return
}

func main() {

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	_, err = ReadCefHeader(&args)
	error_check(err, "Error parsing header")

}
