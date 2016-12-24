package diag

import (
     "testing"
         "github.com/caa-dev-apps/cefmdd_v1/diag"
)

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

var (
    s_diag = diag.NewDiag()
)

func Test_mdd_all_should_fail(t *testing.T) {

    l_test_data := [][]string {
        { "MISSION_REGION",              "Magnetosheath"},
        { "1",                          "FILE_TYPE",                   "cdf"},
        { "",                          "GENERATION_DATE",             "2011-10-09T00:00:00Z" },
        { "2",                          "MISSION_TIME_SPAN",           "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z", "12345" },
        { "",                          "FILE_SIZE",                   "12345" },
        { "2",                          "LOGICAL_FILE_ID",             "" },
        { "",                          "METADATA_VERSION",            "" },
        { "4",                          "SIGNIFICANT_DIGITS",          "12345678"},
    }

    for _, ss := range l_test_data {
        s_diag.Println(ss...)
    }
    
    for _, ss := range l_test_data {
        s_diag.PrintlnN(3, ss...)
    }
    
    s_diag.Printf("%s\t%s\t%s\t%s\t%s\t\n", "1", "2", "3", "4", "5")
    
    s_diag.Flush()
}
