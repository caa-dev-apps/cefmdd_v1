package header

import (
//    "fmt"
    "errors"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

///////////////////////////////////////////////////////////////////////////////
//

func (h *CefHeaderData) getAttrFirstQuoted(key string) (v0 string, err error) {

    vs, err := h.getAttr(key);
    if err == nil {  
        v0 = vs[0]
        if utils.Is_quoted_string(v0) == false {
            err = errors.New(key + " attribute is unquoted string - " + v0)
        } 
    }
        
    return
}


func (h *CefHeaderData) getAttrFirstQuotedTrimed(key string) (v0 string, err error) {
    
    v0, err = h.getAttrFirstQuoted(key)
    if err != nil {
        return
    }    

    v0 = utils.Trim_quoted_string(v0)
    if len(v0) == 0 {
        err = errors.New("Attr: " + key + " length = 0")
    } 
    
    return
}

///////////////////////////////////////////////////////////////////////////////
//

//- Rule: Name of file on disk equals FILE_NAME
func (h *CefHeaderData) check_attr_FILE_NAME() (err error) {
    
    v0, err := h.getAttrFirstQuoted("FILE_NAME")
    if err != nil {
        return
    }
    
    v0 = utils.Trim_quoted_string(v0)
    l_filename := utils.GetCefArgs().GetFilename()

//x fmt.Println("l_filename:", l_filename)
//x fmt.Println("v0:", v0)

    if l_filename != v0 {
        err = errors.New("FILE_NAME attribute(" + v0 + ") - mismatches actual (" + l_filename + ")")
    } 
    
    return
}

//2019smcc
func (h *CefHeaderData) check_attr_DATA_UNTIL()(err error) {
    
    v1, err := h.getAttrFirstQuoted("DATA_UNTIL")
    v1=v1
    if err != nil {
        err = errors.New("DATA_UNTIL required")
        return 
    }

   // diag.Todo("Check_attr_DATA_UNTIL")  
    return
}

func (h *CefHeaderData) check_attr_FILE_FORMAT_VERSION() (err error) {

 //   diag.Todo("Check_attr_FILE_FORMAT_VERSION")  
    
    return
}
        
func (h *CefHeaderData) Attr_Checks() (err error) {
    
	aks := []string{
		"FILE_NAME",
		"FILE_FORMAT_VERSION",
		"DATA_UNTIL",
	}
	for _, k := range aks {
		if v, err := h.getAttr(k); err == nil {
	            diag.Trace(diag.Cyan(" " +k), v)			
		} else {
	            diag.Error("Attr Checks ", err)   
			return err
		}
	}

  
    err = h.check_attr_FILE_NAME()
    h.print_results("FILE_NAME checks", err) 
    h.check_attr_FILE_FORMAT_VERSION()
//    err = h.check_attr_DATA_UNTIL()

    return err
}

