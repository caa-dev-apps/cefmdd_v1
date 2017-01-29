package header

import (
    "fmt"
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


        
func (h *CefHeaderData) Var_Checks() (err error) {
    h.print_results("VAR CHECKS TODO", err) 
    
    var_list := h.Vars().List()
    var_map  := h.Vars().Map()

    for k, v := range var_list {
        v_map := v.Map()
        var_name := v_map["variable_name"]

        fmt.Println("\n\n")
        fmt.Println("k:", k)
        fmt.Println("n:", var_name)
//.         fmt.Println(v)


        depend0, p := v_map["DEPEND_0"]
        if p == true {
            fmt.Println("DEPEND_0 = ", depend0)

            v1_map, p1 := var_map[depend0[0]]
            if p1 == false {
                diag.Error("DEPEND_0 VAR not found", depend0[0])
            } else {
                fmt.Println("DEPEND_0 VAR OK", depend0[0], v1_map.Map()["variable_name"])
            }
        }

    }    

    return
}

//x h.dumpMeta()













