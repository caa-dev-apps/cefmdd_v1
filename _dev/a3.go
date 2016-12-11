package main

import (
	"fmt"
    "strings"
)


//x func main() {
//x     
//x         s1 := "abcdefghijklmnopq_v01.cef"
//x         ix := strings.LastIndex(s1, ".")
//x         
//x         if ix >= 0 {
//x             fmt.Println(s1)
//x             fmt.Println(s1[0:ix])
//x             
//x             fn := s1[0:ix]
//x 
//x             ok := strings.HasSuffix(fn, "v01") 
//x             
//x             fmt.Println("Has suffix = ", ok)
//x         }
//x         
//x }


func main() {
    
        s1 := "C3_CP_EDI_QZC__20111021_V01.cef"
        ix := strings.LastIndex(s1, ".")
        
        if ix >= 0 {
            fmt.Println(s1)
            fmt.Println(s1[0:ix])
            
            fn := s1[0:ix]

            ok := strings.HasSuffix(fn, "01") 
            
            fmt.Println("Has suffix = ", ok)
        }
        
}
