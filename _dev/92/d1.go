package main

import(
    "fmt"
    "regexp"
    "time"
//x     "strings"
    )

    
    
    
    
    
var (
    REGX_REPRESENTATION_, _ := regexp.Compile(`REPRESENTATION_([\d]+)`)
    REGX_LABEL_, _ := regexp.Compile(`LABEL_([\d]+)`)
    REGX_DEPEND_, _ := regexp.Compile(`DEPEND_([\d]+)`)
)

func (h *CefHeaderData) check_mdd(kv *KeyVal) (err error) {

    switch {
        case kv.key == "ENTRY": return
        case kv.key == "FILLVAL": return        // multiple types - depends on var type
        case kv.key == "SIZES": return          // type is FORMAT can be 1 or 1,2 for e.g.
        default:
    }    
    
    
    
    
func main() {
    fmt.Printf("Hello, %s!\n", "World")
    
    r, _ := regexp.Compile(`REPRESENTATION_([\d]+)`)
    l, _ := regexp.Compile(`LABEL_([\d]+)`)
    d, _ := regexp.Compile(`DEPEND_([\d]+)`)
    
    
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
    
    for i, t := range ts {
        fmt.Println(i, t, "->", r.MatchString(t))        
        fmt.Println(i, t, "->", l.MatchString(t))        
        fmt.Println(i, t, "->", d.MatchString(t))        
    }
    
    
    start := time.Now()
    
    c := 0
    for i:=0; i<1000; i++ {
        //x //x s2 := s1 + string(i+48)
        s2 := fmt.Sprintf(`REPRESENTATION_%d`, i)
        
        //x r, _ := regexp.Compile(`REPRESENTATION_([\d]+)`)
        if r.MatchString(s2) {
            c++
        }
        
//x         if strings.HasPrefix(s2, `REPRESENTATION_`) {
//x             c++
//x         }
        
        //x fmt.Println(s2)
    }
    
    fmt.Printf("Count => %d \n", c)
    elapsed := time.Since(start)
    fmt.Printf("Elapsed time %s", elapsed)
    
}    
    