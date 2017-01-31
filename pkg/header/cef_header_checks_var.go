package header

import (
    "fmt"
    "regexp"
    "strconv"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
//x     "errors"
//x     "strings"
//x     "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
//x     "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

//.             type Attrs struct {
//.                 m_map map[string][]string 
//.             }
//. 
//.             func NewAttrs() *Attrs {
//.                 return &Attrs{m_map: make(map[string][]string)}
//.             }
//. 
//.             func (a *Attrs) Map() (map[string][]string) {
//.                 return a.m_map
//.             }
//. 
//.             func (m *Attrs) push_kv(kv *readers.KeyVal) (err error) {
//. 
//.                 val, is_present := m.m_map[kv.Key]
//.                 if is_present == true {
//.                     val = append(val, kv.Val...)
//.                 } else {
//.                     val = kv.Val
//.                 }
//.                 
//.                 m.m_map[kv.Key] = val
//.                 
//.                 return err
//.             }
//. 
//.             func (m *Attrs) push_k_v(k, v string) (err error) {
//.                 vs := []string{ v }
//. 
//.                 kv := readers.NewMetaKeyVal(k , vs)
//. 
//.                 return m.push_kv(kv)
//.             }

//.             type Vars struct {
//.                 m_map map[string]Attrs
//.                 m_list []Attrs
//.             }
//. 
//.             func NewVars() *Vars {
//.                 return &Vars{
//.                     m_map: make(map[string]Attrs),
//.                     m_list: make([]Attrs, 0),
//.                 }
//.             }
//. 
//.             func (v *Vars) Map() (map[string]Attrs) {
//.                 return v.m_map
//.             }
//. 
//.             func (v *Vars) List() ([]Attrs) {
//.                 return v.m_list
//.             }
//. 
//.             type CefHeaderData struct {
//.                 m_state State
//.                 m_name  string
//.                 m_cur* Attrs
//.                 
//.                 m_attrs* Attrs
//.                 m_meta* Meta
//.                 m_vars* Vars
//.             }

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
//x //.         fmt.Println(v)
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

func (h *CefHeaderData) Depend_N_Checks(a Attrs) (err error) {
    // all vars
    vars_map  := h.Vars().Map()

    a_map := a.Map()
    var_name := a_map["variable_name"]
    diag.Info("n:", var_name)

    for k, v := range a_map {
        fmt.Printf("k: %-30s v: %v\n", k, v)

        // regexp.MustCompile("^DEPEND_([0-9]+)$")
        if r := regex_depend_n.FindStringSubmatch(k); r != nil {
            if i, err := strconv.Atoi(r[1]); err == nil {
                vd_map, p1 := vars_map[v[0]]
                if p1 == false {
                    diag.Error("DEPEND_TEST: ref var not found", k, v[0])
                } else if i > 0 {
                    // need to check sizes

                    diag.Todo("DEPENDS_N > 0 - check sizes", vd_map.Map()["variable_name"])
                } else {
                    diag.Todo("Still working on this!")
                }
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

        h.Depend_N_Checks(v)
    }    

    return
}

//x h.dumpMeta()












