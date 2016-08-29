package main

import (
    "strings"
    "errors"
    "fmt"
    "encoding/json"
    "os"
    "strconv"
    
//x 	"io/ioutil"
 	"log"
//x 	"net/http"    
)

//          Current State	                Key In	            Val In	            Checks	                            Output function	            Next State
//          ATTR                            K	                V	                Well known k	                    Data.File.k = v	            A
//                                          K	                V	                Unexpected		                                                Error
//                                          META START	        NAME	            First Occurrence + Name             Expected Data.M[n] = v	    M
//                                                                                  Unexpected		                                                Error
//                                          VAR START	        NAME	            First Occurrence + Name Expected	Data.V[n] = v	            V
//                                                                                  Unexpected		                                                Error
//--------------------------------------------------------------------------------------------------------------------------------------------------------------
//          META                            K	                V	                Key Expected	                    Data.M[k] = v	            M
//                                          K	                V	                Key Unexpected		                                            Error
//                                          META START				                                                                                Error
//                                          VAR START				                                                                                Error
//                                          VAR END				                                                                                    Error
//                                          META_END	        V	                Name Matches start		                                        A
//                                          META_END	        V	                Name mismatch		                                            Error
//--------------------------------------------------------------------------------------------------------------------------------------------------------------
//          VAR                             K	                V	                Key Expected	                    Data.M[k] = v	            V
//                                          K	                V	                Key Unexpected		                                            Error
//                                          VAR START				                                                                                Error
//                                          META START				                                                                                Error
//                                          META END				                                                                                Error
//                                          VAR_END	            V	                Name Matches start		                                        A
//                                          VAR_END	            V	                Name mismatch		                                            Error

///////////////////////////////////////////////////////////////////////////////
//

type State int

const (
	ATTR State = iota
	META
	VAR
	ERROR
)

// used for Global Attributes + Meta and Var  (see KeyVal string -> []string)
type Attrs struct {
    m_map map[string][]string 
}

func NewAttrs() *Attrs {
    return &Attrs{m_map: make(map[string][]string)}
}

func (m *Attrs) push_kv(kv *KeyVal) (err error) {

    val, is_present := m.m_map[kv.key]
    if is_present == true {
        val = append(val, kv.val...)
    } else {
        val = kv.val
    }
    
    m.m_map[kv.key] = val
    
	return err
}


type Meta struct {
    m_map map[string]Attrs
}

func NewMeta() *Meta {
    return &Meta{m_map: make(map[string]Attrs)}
}



type Vars struct {
    m_map map[string]Attrs
    m_list []Attrs
}

func NewVars() *Vars {
    return &Vars{
        m_map: make(map[string]Attrs),
        m_list: make([]Attrs, 0),
    }
}


type CefHeaderData struct {
	m_state State
	m_name  string
    m_cur* Attrs
    
    m_attrs* Attrs
    m_meta* Meta
    m_vars* Vars
}

func NewCefHeaderData() (h *CefHeaderData) {
    
    h = &CefHeaderData{ m_attrs: NewAttrs(),
                        m_meta: NewMeta(),
                        m_vars: NewVars() }    
                        
    h.init()
    return
}

func (h *CefHeaderData) init()  {
    h.m_state = ATTR
    h.m_cur = h.m_attrs
    
	return
}


func dumpJSON(s string, v interface{}) (err error) {
    
	j1, err := json.Marshal(v)
	//x j1, err := json.MarshalIndent(v, "-", "\t")
	if err != nil {
		fmt.Println("error:", err)
        return
	}
    fmt.Println(s)
	os.Stdout.Write(j1)    
    fmt.Println("")
    
    return
}

func dumpJSONMap(s string, m map[string]Attrs) (err error) {
    for k, v := range m {
        
        err = dumpJSON(s+k, v.m_map)
        if err != nil {
            return
        }
    }
    
    return
}

func dumpJSONList(s string, l []Attrs) (err error) {
    for k, v := range l {
        
        err = dumpJSON(s+strconv.Itoa(k), v.m_map)
        if err != nil {
            return
        }
    }
    
    return
}


func (h *CefHeaderData) dump() (err error) {

    fmt.Println("--\n", h.m_attrs)
    fmt.Println("--\n", h.m_meta)
    fmt.Println("--\n", h.m_vars)
    
    
    fmt.Println("\n")
    dumpJSON("Attrs:", h.m_attrs.m_map)
    fmt.Println("\n")
    dumpJSONMap("Meta:", h.m_meta.m_map)
    fmt.Println("\n")
    dumpJSONMap("Var:", h.m_vars.m_map)
    fmt.Println("\n")
    dumpJSONList("Var:", h.m_vars.m_list)
    fmt.Println("\n")
    
	return err
}

//x func mooi_log(s1, s2 string) {
//x     fmt.Println(s1, s2)
//x }

///////////////////////////////////////////////////////////////////////////////
//

var (
    s_mdd_data = NewMddData()
)

//i type KeyVal struct {
//i 	key string
//i 	val []string
//i     
//i     line Line       // 
//i }

func (h *CefHeaderData) check_mdd(kv *KeyVal) (err error) {

    switch {
        case kv.key == "ENTRY": return
        case kv.key == "FILLVAL": return        // multiple types - depends on var type
        case kv.key == "SIZES": return          // type is FORMAT can be 1 or 1,2 for e.g.
        default:
    }

    err = s_mdd_data.test_input(kv)
    if err != nil {
        //x log.Print(err)        
        log.Printf("%s  (Key = %s)", err.Error(), kv.key)        
        //x log.Print(err)        
    }   

    return
}

func (h *CefHeaderData) add_kv(kv *KeyVal) (err error) {

	switch {
	case strings.EqualFold("START_META", kv.key) == true:
		mooi_log("START_META", kv.val[0])

		switch h.m_state {
		case ATTR:
            h.m_state = META
            h.m_name = kv.val[0]
            h.m_cur = NewAttrs()
            
		default:
			return errors.New("START_META: invalid State")
		}

	case strings.EqualFold("START_VARIABLE", kv.key) == true:
		mooi_log("START_VARIABLE", kv.val[0])

		switch h.m_state {
		case ATTR:
            h.m_state = VAR
            h.m_name = kv.val[0]
            h.m_cur = NewAttrs()            
            
		default:
			return errors.New("START_VARIABLE: invalid State")
		}

	case strings.EqualFold("END_META", kv.key) == true:
		mooi_log("END_META", kv.val[0])

		switch h.m_state {
		case META:
			if h.m_name != kv.val[0] {
				return errors.New("END_META: invalid Name")
			}
            
            h.m_meta.m_map[h.m_name] = *h.m_cur
            h.init()            
            
		default:
			return errors.New("END_META: invalid State")
		}

	case strings.EqualFold("END_VARIABLE", kv.key) == true:
		mooi_log("END_VARIABLE", kv.val[0])

		switch h.m_state {
		case VAR:
			if h.m_name != kv.val[0] {
				return errors.New("END_VARIABLE: invalid Name")
			}
            
            h.m_vars.m_list = append(h.m_vars.m_list, *h.m_cur)
            h.m_vars.m_map[h.m_name] = *h.m_cur
            
            h.init()
                        
		default:
			return errors.New("END_VARIABLE: invalid State")
		}

	default:
		mooi_log(kv.key, kv.val[0])

        // this is where to place the key - val tests //////////////////////////////////////////////
        h.m_cur.push_kv(kv)
        
        h.check_mdd(kv)
        
//x         err := s_mdd_data.test_input(kv)
//x         if err != nil {
//x             log.Print(err)        
//x         }        
        
	}

	return err
}

func NewKV(k, v string) *KeyVal {
    return &KeyVal { key: k, val: []string{ v}}
}

