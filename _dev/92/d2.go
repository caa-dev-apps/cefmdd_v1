package main

import(
    "fmt"
    "regexp"
//x     "time"
//x "strings"
    )

var (
    REGX_REPRESENTATION_, _ = regexp.Compile(`REPRESENTATION_([\d]+)`)
    REGX_LABEL_, _ = regexp.Compile(`LABEL_([\d]+)`)
    REGX_DEPEND_, _ = regexp.Compile(`DEPEND_([\d]+)`)
)

    
func main() {
    fmt.Printf("Hello, %s!\n", "World")
    
    ts := []string {
        "REPRESENTATION_a",
        "REPRESENTATION_",
        "REPRESENTATION_1",
        "REPRESENTATION_2",
        "REPRESENTATION_3123",
        
        "LABEL_1",
        "LABEL_2",
        "LABEL_3",
        
        "DEPEND_1",
        "DEPEND_1",
        "DEPEND_1",
    }
    
    for _, t := range ts {
        //x fmt.Println(i, t, "->", r.MatchString(t))        
        //x fmt.Println(i, t, "->", l.MatchString(t))        
        //x fmt.Println(i, t, "->", d.MatchString(t))        
        
        switch {
            case REGX_REPRESENTATION_.MatchString(t): fmt.Println("Match ->", t)
            case REGX_LABEL_.MatchString(t): fmt.Println("Match ->", t)
            case REGX_DEPEND_.MatchString(t): fmt.Println("Match ->", t)
            default: fmt.Println("NO Match ->", t)
        }    
    }
    
//x     start := time.Now()
//x     
//x     c := 0
//x     for i:=0; i<1000; i++ {
//x         //x //x s2 := s1 + string(i+48)
//x         s2 := fmt.Sprintf(`REPRESENTATION_%d`, i)
//x         
//x         if r.MatchString(s2) {
//x             c++
//x         }
//x     }
//x     
//x     fmt.Printf("Count => %d \n", c)
//x     elapsed := time.Since(start)
//x     fmt.Printf("Elapsed time %s", elapsed)
    
}    
    