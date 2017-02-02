package header

import (
    "fmt"
    "regexp"
    "errors"
    "strconv"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

//x             type Attrs struct {
//x                 m_map map[string][]string 
//x             }
//x 
//x             func NewAttrs() *Attrs {
//x                 return &Attrs{m_map: make(map[string][]string)}
//x             }
//x 
//x             func (a *Attrs) Map() (map[string][]string) {
//x                 return a.m_map
//x             }
//x 
//x             func (m *Attrs) push_kv(kv *readers.KeyVal) (err error) {
//x 
//x                 val, is_present := m.m_map[kv.Key]
//x                 if is_present == true {
//x                     val = append(val, kv.Val...)
//x                 } else {
//x                     val = kv.Val
//x                 }
//x                 
//x                 m.m_map[kv.Key] = val
//x                 
//x                 return err
//x             }
//x 
//x             func (m *Attrs) push_k_v(k, v string) (err error) {
//x                 vs := []string{ v }
//x 
//x                 kv := readers.NewMetaKeyVal(k , vs)
//x 
//x                 return m.push_kv(kv)
//x             }

//x             type Vars struct {
//x                 m_map map[string]Attrs
//x                 m_list []Attrs
//x             }
//x 
//x             func NewVars() *Vars {
//x                 return &Vars{
//x                     m_map: make(map[string]Attrs),
//x                     m_list: make([]Attrs, 0),
//x                 }
//x             }
//x 
//x             func (v *Vars) Map() (map[string]Attrs) {
//x                 return v.m_map
//x             }
//x 
//x             func (v *Vars) List() ([]Attrs) {
//x                 return v.m_list
//x             }
//x 
//x             type CefHeaderData struct {
//x                 m_state State
//x                 m_name  string
//x                 m_cur* Attrs
//x                 
//x                 m_attrs* Attrs
//x                 m_meta* Meta
//x                 m_vars* Vars
//x             }

//x func (h *CefHeaderData) Var_Checks() (err error) {
//x     h.print_results("VAR CHECKS TODO", err) 
//x     
//x     var_list := h.Vars().List()
//x     var_map  := h.Vars().Map()
//x 
//x     for k, v := range var_list {
//x         v_map := v.Map()
//x         var_name := v_map["variable_name"]
//x 
//x         fmt.Println("\n\n")
//x         fmt.Println("k:", k)
//x         fmt.Println("n:", var_name)
//x //x         fmt.Println(v)
//x 
//x 
//x         depend0, p := v_map["DEPEND_0"]
//x         if p == true {
//x             fmt.Println("DEPEND_0 = ", depend0)
//x 
//x             v1_map, p1 := var_map[depend0[0]]
//x             if p1 == false {
//x                 diag.Error("DEPEND_0 VAR not found", depend0[0])
//x             } else {
//x                 fmt.Println("DEPEND_0 VAR OK", depend0[0], v1_map.Map()["variable_name"])
//x             }
//x         }
//x 
//x     }    
//x 
//x     return
//x }

var (
    regex_depend_n *regexp.Regexp   
)

func init() {
    regex_depend_n = regexp.MustCompile("^DEPEND_([0-9]+)$")
}


// Checks VALUE_TYPE and SIZES when DEPEND_N N > 0
func (h *CefHeaderData) Depend_N_Checks(a1 Attrs) (err error) {
    // all vars
    a1_map := a1.Map()
    vars_map := h.Vars().Map()

    var_name := a1_map["variable_name"]
    //x diag.Info("n:", var_name)

    for k, v := range a1_map {
        //x fmt.Printf("k: %-30s v: %v\n", k, v)

        // regexp.MustCompile("^DEPEND_([0-9]+)$")
        r := regex_depend_n.FindStringSubmatch(k)
        if r != nil {
            i, err := strconv.Atoi(r[1])
            if  err == nil {
                // var pointed to by DEPEND
                a2, p0 := vars_map[v[0]]
                if p0 == false {
                    diag.Error("DEPEND_TEST: ref var not found", k, v[0])
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

                //x diag.Trace(diag.BoldGreen("ok DEPEND checks(VALUE_TYPE,SIZES)"), "variable:", var_name, k, v[0])
                diag.Trace(diag.BoldGreen("ok"), k, v[0], "variable:", var_name, )
            } else {
                return errors.New(fmt.Sprintf("Malformed: %v %v", k, v))
            }
        }    
    }

    return
}

func (h *CefHeaderData) Var_Checks() (err error) {
    h.print_results("VAR CHECKS TODO", err) 
    
    var_list := h.Vars().List()

    for k, v := range var_list {
        fmt.Println("k:", k)

        err = h.Depend_N_Checks(v)
        if err != nil {
            return
        }
    }    

    return
}













