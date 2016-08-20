package main

import (
	"fmt"
	"os"
)

func error_check(err error, i_s string) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error-Check: ", i_s)
		os.Exit(1)
	}
}

func fileExists(name string) (isReq bool, err error) {

	info, err := os.Stat(name)

	if err == nil {
		isReq = info.Mode().IsRegular()
	}

	return
}
