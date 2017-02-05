package header

import (
    "fmt"
    "regexp"
    "errors"
    "strconv"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
)

//-             type Attrs struct {
//-                 m_map map[string][]string 
//-             }
//- 
//-             func NewAttrs() *Attrs {
//-                 return &Attrs{m_map: make(map[string][]string)}
//-             }
//- 
//-             func (a *Attrs) Map() (map[string][]string) {
//-                 return a.m_map
//-             }
//- 
//-             func (m *Attrs) push_kv(kv *readers.KeyVal) (err error) {
//- 
//-                 val, is_present := m.m_map[kv.Key]
//-                 if is_present == true {
//-                     val = append(val, kv.Val...)
//-                 } else {
//-                     val = kv.Val
//-                 }
//-                 
//-                 m.m_map[kv.Key] = val
//-                 
//-                 return err
//-             }
//- 
//-             func (m *Attrs) push_k_v(k, v string) (err error) {
//-                 vs := []string{ v }
//- 
//-                 kv := readers.NewMetaKeyVal(k , vs)
//- 
//-                 return m.push_kv(kv)
//-             }

//-             type Vars struct {
//-                 m_map map[string]Attrs
//-                 m_list []Attrs
//-             }
//- 
//-             func NewVars() *Vars {
//-                 return &Vars{
//-                     m_map: make(map[string]Attrs),
//-                     m_list: make([]Attrs, 0),
//-                 }
//-             }
//- 
//-             func (v *Vars) Map() (map[string]Attrs) {
//-                 return v.m_map
//-             }
//- 
//-             func (v *Vars) List() ([]Attrs) {
//-                 return v.m_list
//-             }
//- 
//-             type CefHeaderData struct {
//-                 m_state State
//-                 m_name  string
//-                 m_cur* Attrs
//-                 
//-                 m_attrs* Attrs
//-                 m_meta* Meta
//-                 m_vars* Vars
//-             }

var (
    regex_depend_n *regexp.Regexp   
)

func init() {
    regex_depend_n = regexp.MustCompile("^DEPEND_([0-9]+)$")
}

func (h *CefHeaderData) Depend_N_Checks(a1 Attrs) (err error) {
    // all vars
    a1_map := a1.Map()
    vars_map := h.Vars().Map()

    var_name := a1_map["variable_name"]
    //- diag.Info("n:", var_name)

    for k, v := range a1_map {
        //- fmt.Printf("k: %-30s v: %v\n", k, v)

        // regexp.MustCompile("^DEPEND_([0-9]+)$")
        r := regex_depend_n.FindStringSubmatch(k)
        if r != nil {
            i, err := strconv.Atoi(r[1])
            if  err == nil {
                // var pointed to by DEPEND
                a2, p0 := vars_map[v[0]]
                if p0 == false {
                    //- diag.Error("DEPEND_TEST: ref var not found", k, v[0])
                    return errors.New(fmt.Sprintf("DEPEND_TEST: ref var not found %v %v ", k, v[0]))
                } else if i > 0 {


                    vt1, p1 := a1_map["VALUE_TYPE"]
                    if p1 == false {
                        return errors.New(fmt.Sprintf("Missing VALUE_TYPE: variable (%s) ", var_name))
                    }

                    sz1, p1 := a1_map["SIZES"]
                    if p1 == false {
                        return errors.New(fmt.Sprintf("Missing SIZES: variable (%s) ", var_name))
                    }


                    if len(sz1) < i {
                        return errors.New(fmt.Sprintf("DEPEND index(%d) out of range for SIZES, variable (%s) : %v %v %v", i, var_name, sz1, k, v))
                    }

                    a2_map:= a2.Map()

                    vt2, p2 := a2_map["VALUE_TYPE"]
                    if p2 == false {
                        return errors.New(fmt.Sprintf("Missing VALUE_TYPE for DEPEND variable: %v %v", k, v))
                    }


                    sz2, p2 := a2_map["SIZES"]
                    if p2 == false {
                        return errors.New(fmt.Sprintf("Missing SIZES for DEPEND variable: %v %v", k, v))
                    }


                    // value_type checks
                    if vt1[0] != vt2[0] {
                        return errors.New(fmt.Sprintf("Missmatch VALUE_TYPE: %v %v %v %v", vt1, vt2, k, v))
                    }

                    if sz1[i-1] != sz2[0] {
                        return errors.New(fmt.Sprintf("Missmatch SIZES for DEPEND Variable, index(%d) : %v %v %v %v", i-1, sz1, sz2, k, v))
                    }
                } 

                diag.Trace(diag.BoldGreen(" ok:"), k, v[0], "variable:", var_name)
            } else {
                return errors.New(fmt.Sprintf("Malformed: %v %v", k, v))
            }
        }    
    }

    return
}

// Checks DELTA_PLUS, DELTA_MINUS, ERROR_PLUS, ERROR_MINUS, QUALITY
func (h *CefHeaderData) Numeric_Or_Variable(k string, a1 Attrs) (err error) {
    // all vars
    a1_map := a1.Map()
    vars_map := h.Vars().Map()

    // only need to test if present
    v, p := a1_map[k]
    if p == true {

        v0 := utils.Trim_quoted_string(v[0])
        if utils.Is_Numerical(v0) == true {
            diag.Trace(diag.BoldGreen(" ok:"), k, v0)
            return
        }

        // check if it's a variable
        _, p0 := vars_map[v0]
        if p0 == false {
            return errors.New(fmt.Sprintf("%s: ref var not found %v  %v", k, v0, v))
        } else {
            diag.Trace(diag.BoldGreen(" ok:"), k, v0)
        }
    }

    return
}

// Checks DELTA_PLUS, DELTA_MINUS, ERROR_PLUS, ERROR_MINUS, QUALITY
func (h *CefHeaderData) NumerIc_or_Variable_Checks(a1 Attrs) (err error) {

    keys := []string {
        "DELTA_PLUS",
        "DELTA_MINUS",
        "ERROR_MINUS",
        "ERROR_PLUS",
        "QUALITY",
        "MIN_TIME_RESOLUTION",
        "MAX_TIME_RESOLUTION",
    }

    for _, k := range keys {
        err = h.Numeric_Or_Variable(k, a1) 
        if err != nil {
            return
        }
    }

    return
}

func (h *CefHeaderData) Var_Checks() (err error) {
    err_count := 0 
    var_list := h.Vars().List()

    for ix, v := range var_list {
        
        diag.Trace(diag.Cyan(fmt.Sprintf(" %d %s", ix, v.Map()["variable_name"])))

        err = h.Depend_N_Checks(v)
        if err != nil {
            diag.Error("Depend Checks", err)
            err_count++
        }

        err = h.NumerIc_or_Variable_Checks(v)
        if err != nil {
            diag.Error("Plus/Minus/Quality/etc Checks", err)
            err_count++
        }
    }    

    if err_count > 0 {
        return errors.New(fmt.Sprintf("Var Checks %d error(s) ", err_count))
    }

    return
}
