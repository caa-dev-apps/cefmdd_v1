package main

import (
    "strings"
    "errors"
    "fmt"
    "encoding/json"
    "os"
    "strconv"
    "regexp"
//x 	"io/ioutil"
//x  	"log"
//x     "github.com/fatih/color"
//x 	"net/http"    
    "github.com/caa-dev-apps/cefmdd_v1/readers"
    "github.com/caa-dev-apps/cefmdd_v1/rules"
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

// used for Global Attributes + Meta and Var  (see readers.KeyVal string -> []string)
type Attrs struct {
    m_map map[string][]string 
}

func NewAttrs() *Attrs {
    return &Attrs{m_map: make(map[string][]string)}
}

func (m *Attrs) push_kv(kv *readers.KeyVal) (err error) {

    val, is_present := m.m_map[kv.Key]
    if is_present == true {
        val = append(val, kv.Val...)
    } else {
        val = kv.Val
    }
    
    m.m_map[kv.Key] = val
    
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


func (h *CefHeaderData) dumpAttrs()  {
    dumpJSON("Attrs:", h.m_attrs.m_map)
    fmt.Println("\n")
}

func (h *CefHeaderData) dumpMeta()  {
    dumpJSONMap("Meta:", h.m_meta.m_map)
    fmt.Println("\n")
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

func (h *CefHeaderData) getAttr(k string) (v []string, err error) {
    
    v, present := h.m_attrs.m_map[k]
    if present == false {
        err = errors.New("Error: keyword not presnt " + k)
    } 
    
    return
}


func (h *CefHeaderData) getMeta(k string) (entry, value_type []string, err error) {
    
    v, present := h.m_meta.m_map[k]
    if present == false {
        err = errors.New("Error: meta not presnt " + k)
    } else {
        entry, present = v.m_map[`ENTRY`]
        if present == false {
            err = errors.New("Error: meta:ENTRY not presnt " + k)
        } else {
        }
        
        value_type, _ = v.m_map[`VALUE_TYPE`]
    }
    
    return
}

func (h *CefHeaderData) getMetaEntry(k string) (entry []string, err error) {
    
    entry, _, err = h.getMeta(k)
    
    return
}


///////////////////////////////////////////////////////////////////////////////
//
 
var (
    s_mdd_data = rules.NewMddData()
)

var (
    REGX_REPRESENTATION_i, _ = regexp.Compile(`REPRESENTATION_([\d]+)`)
    REGX_LABEL_i, _ = regexp.Compile(`LABEL_([\d]+)`)
    REGX_DEPEND_i, _ = regexp.Compile(`DEPEND_([\d]+)`)
    
    s_diag = NewDiag()
)



func (h *CefHeaderData) check_mdd(kv *readers.KeyVal) (err error) {

    l_kv := kv

    switch {
        case kv.Key == "ENTRY":     fmt.Println(BoldBlue("delay-meta-check-", "ENTRY",   " ",  kv.Val));  return
        case kv.Key == "FILLVAL":   fmt.Println(BoldBlue("delay-check-", "FILLVAL", " ",  kv.Val));  return        // multiple types - depends on var type
        case kv.Key == "SIZES":     fmt.Println(BoldBlue("delay-check-", "SIZES",   " ",  kv.Val));  return        // type is FORMAT can be 1 or 1,2 for e.g.
        
        case REGX_REPRESENTATION_i.MatchString(kv.Key):  l_kv = kv.NewSwitchKey(`REPRESENTATION_i`)
        case REGX_LABEL_i.MatchString(kv.Key):           l_kv = kv.NewSwitchKey(`LABEL_i`)
        case REGX_DEPEND_i.MatchString(kv.Key):          l_kv = kv.NewSwitchKey(`DEPEND_i`)
        
        default:
    }

    err = s_mdd_data.Test_input(l_kv)
    if err != nil {
        fmt.Println(BoldRed(err.Error()))
    } else {
        fmt.Println(BoldGreen("Ok-readers.KeyVal"), kv.Key, kv.Val)
    }

    return
}


func (h *CefHeaderData) check_mdd_meta_etx() (err error) {
    
    l_kv := readers.NewMetaKeyVal(h.m_name, h.m_cur.m_map[`ENTRY`])
    
    err = s_mdd_data.Test_input(l_kv)
    if err != nil {
        fmt.Println(BoldRed(err.Error()))
        // log.Printf("\t\tTESTING-2 %s %s\n",h.m_name, h.m_cur.m_map[`ENTRY`])
    } else {
        fmt.Println(BoldCyan("ok-meta"), h.m_name)
    }
    
    
    return
}

func (h *CefHeaderData) add_kv(kv *readers.KeyVal) (err error) {

	switch {
	case strings.EqualFold("START_META", kv.Key) == true:
		//d mooi_log("START_META", kv.Val[0])

		switch h.m_state {
		case ATTR:
            h.m_state = META
            h.m_name = kv.Val[0]
            h.m_cur = NewAttrs()
            
		default:
			return errors.New("START_META: invalid State")
		}

	case strings.EqualFold("START_VARIABLE", kv.Key) == true:
		//d mooi_log("START_VARIABLE", kv.Val[0])

		switch h.m_state {
		case ATTR:
            h.m_state = VAR
            h.m_name = kv.Val[0]
            h.m_cur = NewAttrs()            
            
		default:
			return errors.New("START_VARIABLE: invalid State")
		}

	case strings.EqualFold("END_META", kv.Key) == true:
		//d mooi_log("END_META", kv.Val[0])

		switch h.m_state {
		case META:
			if h.m_name != kv.Val[0] {
				return errors.New("END_META: invalid Name")
			}
            
            h.m_meta.m_map[h.m_name] = *h.m_cur
            
            h.check_mdd_meta_etx()
            
            h.init()            
            
		default:
			return errors.New("END_META: invalid State")
		}

	case strings.EqualFold("END_VARIABLE", kv.Key) == true:
		//d mooi_log("END_VARIABLE", kv.Val[0])

		switch h.m_state {
		case VAR:
			if h.m_name != kv.Val[0] {
				return errors.New("END_VARIABLE: invalid Name")
			}
            
            h.m_vars.m_list = append(h.m_vars.m_list, *h.m_cur)
            h.m_vars.m_map[h.m_name] = *h.m_cur
            
            
            
            h.init()
                        
		default:
			return errors.New("END_VARIABLE: invalid State")
		}

	default:
		//d mooi_log(kv.Key, kv.Val[0])

        // this is where to place the key - val tests //////////////////////////////////////////////
        h.m_cur.push_kv(kv)
        
        h.check_mdd(kv)
        
    
        
	}

	return err
}

func NewKV(k, v string) *readers.KeyVal {
    return &readers.KeyVal { Key: k, Val: []string{ v}}
}

