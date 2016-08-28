package main

import (
	//x "log"
	"fmt"
	"testing"
)

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

var (
	m_path_eums     = "C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Enums.csv"
	m_path_keywords = "C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Tables.csv"
)

func Test_all_should_fail(t *testing.T) {

	for r := range read_csv(m_path_eums) {
//x 		fmt.Println("0: ", r[0], "1: ", r[1], "2: ", r[2], "3: ", r[3])
		fmt.Println("0: ", r[0], "1: ", r[1])
	}
}
