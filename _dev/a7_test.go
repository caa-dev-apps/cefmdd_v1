package main

//go -v test a7_test.go a7.go

import (
    "testing"
	"fmt"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
//

func array_line_reader(i_lines []string) chan string {
    output := make(chan string, 16)

    go func() {
        defer close(output)

        for ix, l := range i_lines {
            fmt.Println(ix, l)

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

    for l := range rowTokeniser(l_reader, i_len) {
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

    for l := range rowTokeniser(l_reader, i_len) {
        if l.Err != nil {
            t.Error(l.Err, "\n", l)
        } else {
            t.Log(l)
        }
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////
//

//1 func TestReader_000(t *testing.T) {
//1
//1    l_test_data := []string {
//1//x         "!                                            D1                             D2", 
//1//x         "!        time              Stat  pa90 lat_gse lon_gse  phi thet  pa90 lat_gse lon_gse  phi thet  ", 
//1        "2011-10-21T14:24:16.648437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:16.773437Z,  1,    0, -39.41,  38.65,  45, 103,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:16.898437Z,  1,    0, -69.68,  36.49,  55, 114,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.023437Z,  1,    0, -79.93,-135.84,  66, 115,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.148437Z,  1,    0, -49.76,-141.11,  76, 106,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.273437Z,  1,    0, -19.62,-141.73,  86,  89,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.398437Z,  1,    0,  10.51,-142.97,  99,  78,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.523437Z,  1,    0,  40.68,-144.37, 111,  83,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.648437Z,  1,    0,  71.00,-146.79, 121, 107,   -1, 999.99, 999.99, 999, 999", 
//1//x         "2011-10-21T14:24:17.773437Z,  1,    0, -79.73,-142.69,  62, 115,   -1, 999.99, 999.99, 999, 999", 
//1//x         "20 11-10-21T14:24:17.898437Z,  1,    0, -49.67,-146.82,  65,  75,   -1, 999.99, 999.99, 999, 999",
//1    }
//1
//1
//1    shouldPass("TestReader_001",
//1                l_test_data,
//1                11,
//1                t)
//1}

//2  func TestReader_001(t *testing.T) {
//2  
//2      l_test_data := []string {
//2          "2011-10-21T14:24:16.648437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
//2      }
//2  
//2      shouldPass("TestReader_001",
//2                  l_test_data,
//2                  11,
//2                  t)
//2  }
//2  
//2  func TestReader_002(t *testing.T) {
//2  
//2      l_test_data := []string {
//2          //x                     x
//2          "2011-10-21T14:24:16.648 437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
//2      }
//2  
//2      shouldFail("TestReader_002",
//2                  l_test_data,
//2                  11,
//2                  t)
//2  }
//2 
//2 func TestReader_003(t *testing.T) {
//2 
//2     l_test_data := []string {
//2         //x                     x
//2         " , 2011-10-21T14:24:16.648 437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
//2     }
//2 
//2     shouldFail("TestReader_002",
//2                 l_test_data,
//2                 11,
//2                 t)
//2 }

func TestReader_004(t *testing.T) {

    l_test_data := []string {
        " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -1000000000, -1000000000, -1000000000, -1000000000 $ ", 
    }

    shouldPass("TestReader_002",
                l_test_data,
                10,
                t)
}

