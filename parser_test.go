package main

import (
     "testing"
)

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

var (
    s_hd = NewCefHeaderData()
)


func Test_header_add_kv(t *testing.T) {

                         
    s_hd.add_kv(NewKV("a1", "a1_s_hd"))
    s_hd.add_kv(NewKV("a2", "a2_s_hd"))
    
    s_hd.add_kv(NewKV("START_META", "META_001"))
        s_hd.add_kv(NewKV("m1a", "m1a_s_hd"))
        s_hd.add_kv(NewKV("m1b", "m1b_s_hd"))
    s_hd.add_kv(NewKV("END_META", "META_001"))
    

    s_hd.add_kv(NewKV("START_META", "META_002"))
        s_hd.add_kv(NewKV("m2a", "m2a_s_hd"))
        s_hd.add_kv(NewKV("m2b", "m2b_s_hd"))
    s_hd.add_kv(NewKV("END_META", "META_002"))


    s_hd.add_kv(NewKV("START_VARIABLE", "VAR_001"))
        s_hd.add_kv(NewKV("s_hda", "s_hda_s_hd"))
        s_hd.add_kv(NewKV("s_hdb", "s_hdb_s_hd"))
        s_hd.add_kv(NewKV("s_hdc", "s_hdc_s_hd"))
    s_hd.add_kv(NewKV("END_VARIABLE", "VAR_001"))
    

    s_hd.add_kv(NewKV("START_VARIABLE", "VAR_002"))
        s_hd.add_kv(NewKV("v2a", "v2a_s_hd"))
        s_hd.add_kv(NewKV("v2b", "v2b_s_hd"))
        s_hd.add_kv(NewKV("v2c", "v2c_s_hd"))
    s_hd.add_kv(NewKV("END_VARIABLE", "VAR_002"))

    s_hd.dump()
}

