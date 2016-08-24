package main

import (
//x     "bufio"
//x     "encoding/csv"
//x     "os"
//x      "errors"
     "fmt"
//x     "io"
     "log"
//x     "strings"
     "testing"
//x     "github.com/caa-dev-apps/cefmdd_v1/_dev"    
)

//- Keywords
//- Key,                    low,    high,   type,               scope,  source,     table,  comments__20160617,,
//- MISSION                 ,1,     1,      Enumerated         ,m,      CAA        ,3,      ,       ,
//- PARENT_MISSION          ,1,     1,      Enumerated         ,,       CAA        ,3,,,
//- MISSION_TIME_SPAN       ,1,     1,      ISO_TIME_RANGE     ,m,      CAA        ,3,,,

//- Enum
//- Keyword,                Definition,     Description,                                            Additional
//- MISSION,                Cluster,        ,
//- MISSION,                DoubleStar,     ,
//- PARENT_MISSION,         Cluster,        Mission or project from which the data was collected:,
//- MISSION_AGENCY,         ESA,            European Space Agency,

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

func Test0001(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    
    fmt.Println("Hello, World!")
    
    mdd_data := NewMddData()
                           
    err := mdd_data.test_input_kv1("LOGICAL_FILE_ID", "SOME_FILE_ID")
    if err != nil {
        log.Print(err)        
    }
    
//x     err = mdd_data.test_input_kv2("LOGICAL_FILE_ID_2", []string{"SOME_FILE_ID"})
//x     if err != nil {
//x         log.Print(err)        
//x     }
    
}



