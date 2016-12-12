package main

import (
	"fmt"
    "strings"
)

var (
    test_data = []string {
        "!                                            D1                             D2", 
        "!        time              Stat  pa90 lat_gse lon_gse  phi thet  pa90 lat_gse lon_gse  phi thet  ", 
        "2011-10-21T14:24:16.648437Z,  1,    0,  -8.38, -24.13, 112,  -1,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:16.773437Z,  1,    0, -39.41,  38.65,  45, 103,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:16.898437Z,  1,    0, -69.68,  36.49,  55, 114,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.023437Z,  1,    0, -79.93,-135.84,  66, 115,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.148437Z,  1,    0, -49.76,-141.11,  76, 106,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.273437Z,  1,    0, -19.62,-141.73,  86,  89,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.398437Z,  1,    0,  10.51,-142.97,  99,  78,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.523437Z,  1,    0,  40.68,-144.37, 111,  83,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.648437Z,  1,    0,  71.00,-146.79, 121, 107,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.773437Z,  1,    0, -79.73,-142.69,  62, 115,   -1, 999.99, 999.99, 999, 999", 
        "2011-10-21T14:24:17.898437Z,  1,    0, -49.67,-146.82,  65,  75,   -1, 999.99, 999.99, 999, 999",
    }
)

//x func EachLine(i_path string) chan Line {
//x     output := make(chan Line, 16)
//x     l_tag := path.Base(i_path)
//x     
//x     go func() {
//x         defer close(output)
//x 
//x         fi, err := os.Open(i_path)
//x         if err != nil {
//x             return
//x         }
//x         defer fi.Close()
//x 
//x         var reader *bufio.Reader
//x         if strings.HasSuffix(i_path, "gz") == false {
//x             reader = bufio.NewReader(fi)
//x         } else {
//x             fz, err := gzip.NewReader(fi)
//x             if err != nil {
//x                 return
//x             }
//x             defer fz.Close()
//x 
//x             reader = bufio.NewReader(fz)
//x         }
//x 
//x         ln  := 0
//x         
//x         for {
//x             line, err := reader.ReadString('\n')
//x             output <- Line{tag: l_tag, ln: ln, line: line }
//x             if err == io.EOF {
//x                 break
//x             }
//x             
//x             ln++
//x         }
//x 
//x 
//x         for ix, l := range test_data {
//x             fmt.Println(ix, l)
//x         }
//x 
//x     }()
//x 
//x     return output
//x }


func line_reader() chan string {
    output := make(chan string, 16)
    
    go func() {
        defer close(output)

        for ix, l := range test_data {
            fmt.Println(ix, l)

            output <- l
        }

    }()

    return output
}


func multi_line_read(n int, eol string) chan []string {
    output := make(chan []string, 16)

    go func() {
        defer close(output)

        cs1 := make([]string, n)
        //x cs1 := []string

        for l := range line_reader() {
            fmt.Println(l)

            cs2 := strings.Split(l, ",")

            if len(cs1) > 0 {
                cs1 = append(cs1, cs2...)
            } else {
                cs1 = cs2
            }

            if len(cs1) >= n {
               output <- cs1 
               cs1 = make([]string, n)
            }
        }

    }()

    return output
}



func read_data() {
    for ls := range multi_line_read(5, "$") {
        fmt.Println(ls)
    }
}

func main() {
    read_data()
}
