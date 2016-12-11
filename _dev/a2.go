package main

import (
	"fmt"
    //x "strings"
)

func is_quoted_string(s string) (r bool) {

    //x if s != nil {
        l := len(s)
        r = l >= 2 && s[0] == '"' && s[l-1] == '"'
    //x }

    fmt.Println("is quoted string -> ", s, r)
    
    return
}


func is_quoted_string2(s *string) (r bool) {

    if s != nil {
        s1 := *s
        l := len(s1)
        r = l >= 2 && s1[0] == '"' && s1[l-1] == '"'
    }

    fmt.Println("is quoted string -> ", *s, r)
    
    return
}


func main() {
    
        s1 := `"a test string"`
        s2 := `"a test stringX`
        s3 := `""`
        
        is_quoted_string2(&s1)
        is_quoted_string2(&s2)
        is_quoted_string2(&s3)
        //x is_quoted_string(nil)
        //x is_quoted_string(1234)
}
