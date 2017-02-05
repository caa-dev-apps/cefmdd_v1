package readers

//go test -v a7_test.go a7.go

import (
    "testing"
	"fmt"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
//

func array_line_reader(i_lines []string) chan string {
    output := make(chan string, 16)

    go func() {
        defer close(output)

        for ix, l := range i_lines {
            diag.Println(ix, l)

            output <- l
        }

    }()

    return output
}

///////////////////////////////////////////////////////////////////////////////////////////////////
//

func shouldFail(i_about string,
                i_test_data []string,
                i_len int,
                t *testing.T) {

    t.Log(i_about)

    l_reader := array_line_reader(i_test_data)

    for l := range DataRecords(l_reader, i_len) {
        if l.Err == nil {
            t.Error("ERROR THIS SHOULD FAIL!!! (but it did NOT) \n", l)
        } else {
            t.Log("FAILED as required: \n", l.Err, "\n", l)
        }
    }
}

func shouldPass(i_about string,
                i_test_data []string,
                i_len int,
                t *testing.T) {

    t.Log(i_about)

    l_reader := array_line_reader(i_test_data)

    for l := range DataRecords(l_reader, i_len) {
        if l.Err != nil {
            t.Error(l.Err, "\n", l)
        } else {
            t.Log(l)
        }
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
//

 func TestReader_002(t *testing.T) {
 
     l_test_data := []string {
         "2011-10-21T14:24:16.648 437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
     }
 
     shouldFail("TestReader_002",
                 l_test_data,
                 11,
                 t)
 }

func TestReader_003(t *testing.T) {

    l_test_data := []string {
        " , 2011-10-21T14:24:16.648 437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
    }

    shouldFail("TestReader_003",
                l_test_data,
                11,
                t)
}

func TestReader_004(t *testing.T) {

    l_test_data := []string {
        " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -1000000000, -1000000000, -1000000000, -1000000000 $ ", 
    }

    shouldPass("TestReader_004",
                l_test_data,
                10,
                t)
}


func TestReader_005(t *testing.T) {

    l_test_data := []string {
        " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -1000000000, -1000000000 ", 
        "  -1000000000, -1000000000 $ ", 
    }

    shouldPass("TestReader_005",
                l_test_data,
                10,
                t)
}

func TestReader_006(t *testing.T) {

    l_test_data := []string {
        " 2002-11-10T04:17:14.758391Z       ", 
        " -502  ,        -596, ",
        "  -1000000000, -1000000000 ",
        " -1.3e10      ,   ",
        "    -763, -1000000000, -1000000000, $", 
        " 2002-11-10T04:17:14.999999Z       ", 
        " -991  ,        -596, ",
          "  -9999999000, -1000000000 ",
        " -1.3e10      ,   ",
        "    -763, -1000000000, -1000000000, $", 
    }

    shouldPass("TestReader_006",
                l_test_data,
                10,
                t)
}

