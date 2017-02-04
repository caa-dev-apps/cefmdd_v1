package header

import (
//x     "fmt"
    "errors"
    "strings"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

///////////////////////////////////////////////////////////////////////////////
//

func (h *CefHeaderData) getMetaEntryFirstQuoted(key string) (v0 string, err error) {

    vs, err := h.getMetaEntry(key);
    if err == nil {  
        v0 = vs[0]
        if utils.Is_quoted_string(v0) == false {
            err = errors.New(key + " META ENTRY is unquoted string - " + v0)
        } 
    }
        
    return
}

func (h *CefHeaderData) getMetaEntryFirstQuotedTrimed(key string) (v0 string, err error) {

    v0, err = h.getMetaEntryFirstQuoted(key)
    if err != nil {
        return
    }    

    v0 = utils.Trim_quoted_string(v0)
    if len(v0) == 0 {
        err = errors.New("Meta:ENTRY: " + key + " length = 0")
    } 
    
    return
}

///////////////////////////////////////////////////////////////////////////////
//

//- START_META     =   DATASET_ID
//-    ENTRY       =   "C3_CP_EDI_QZC"
//- END_META       =   DATASET_ID
//- !
//- START_META     =   LOGICAL_FILE_ID
//-    ENTRY       =   "C3_CP_EDI_QZC__20111021_V01"
//- END_META       =   LOGICAL_FILE_ID

//- Rule: Dataset name matches LOGICAL_FILE_ID (DATASET must be start of logical file id)
func (h *CefHeaderData) check_meta_LOGICAL_FILE_ID() (err error) {
    
    v0_dataset, err := h.getMetaEntryFirstQuoted("DATASET_ID")
    if err != nil {
        return
    }
    
    v0_logical, err := h.getMetaEntryFirstQuoted("LOGICAL_FILE_ID")
    if err != nil {
        return
    }
    
    v0_dataset = utils.Trim_quoted_string(v0_dataset)
    if len(v0_dataset) == 0 {
        err = errors.New("DATASET_ID length = 0")
        return
    } 
    
    v0_logical = utils.Trim_quoted_string(v0_logical)
    if strings.HasPrefix(v0_logical, v0_dataset) == false {
        err = errors.New("DATASET_ID (" + v0_dataset + ") is not a prefix of LOGICAL_FILE_ID (" + v0_logical + ")" )
    }
    
    return
}

//- FILE_NAME="C3_CP_EDI_QZC__20111021_V01.cef"
//- 
//- START_META     = LOGICAL_FILE_ID
//-     ENTRY      = "C3_CP_EDI_QZC__20111021_V01"
//- END_META       = LOGICAL_FILE_ID
//- !
//- START_META     =   FILE_TYPE
//-    ENTRY       =   "cef"
//- END_META       =   FILE_TYPE

//- Filename matches LOGICAL_FILE_ID + FILE_TYPE        
func (h *CefHeaderData) check_meta_FILE_TYPE() (err error) {
    
    v0_filename, err := h.getAttrFirstQuoted("FILE_NAME")
    if err != nil {
        return
    }    
    
    v0_file_type, err := h.getMetaEntryFirstQuoted("FILE_TYPE")
    if err != nil {
        return
    }
    
    v0_logical, err := h.getMetaEntryFirstQuoted("LOGICAL_FILE_ID")
    if err != nil {
        return
    }
    
    v0_filename = utils.Trim_quoted_string(v0_filename)
    if len(v0_filename) == 0 {
        return errors.New("FILE_NAME length = 0")
    } 
    
    v0_file_type = utils.Trim_quoted_string(v0_file_type)
    if len(v0_file_type) == 0 {
        return errors.New("FILE_TYPE length = 0")
    } 
    
    v0_logical = utils.Trim_quoted_string(v0_logical)
    if len(v0_logical) == 0 {
        return errors.New("LOGICAL_FILE_ID length = 0")
    } 
    
    
    //x ix := strings.FirstIndex(v0_filename, ".")
    ix := strings.Index(v0_filename, ".")
    if ix < 0 {
        return errors.New("FILE_NAME does have .cef or .cef.gz suffix")
    }        

    ext0 := v0_filename[ix+1:]
    if strings.HasPrefix(ext0, v0_file_type) == false {
        return errors.New("FILE_TYPE does not follow first '.' in actual filename" + "||" + ext0 + "||" + v0_file_type)
    }
    
    fn := v0_logical + `.` + v0_file_type
    // allow for .gz
    if strings.HasSuffix(v0_filename, fn) == false {
        diag.Println(v0_filename)
        diag.Println(fn)
        
        return errors.New("FILE_NAME does not match LOGICAL_FILE_ID + FILE_TYPE  || " + v0_filename  + " || " + fn)
    }
    
    return
}

//- FILE_TIME_SPAN must be ISO time
//- FIME_TIME_SPAN start time must be before stop time
func (h *CefHeaderData) check_meta_FILE_TIME_SPAN() (err error) {
    
    es, vs, err := h.getMeta(`FILE_TIME_SPAN`)
    
    if err != nil {
        return
    }

    if len(es) == 0 {
        return errors.New("META-ENTRY for FILE_TIME_SPAN Missing")
    }
    
    if len(vs) == 0 {
        return errors.New("META-VALUE_TYPE for FILE_TIME_SPAN Missing")
    }
    
    v1_value_type_FILE_TIME_SPAN := vs[0]
    if "ISO_TIME_RANGE" != v1_value_type_FILE_TIME_SPAN {
        return errors.New("META-VALUE_TYPE for FILE_TIME_SPAN not equal to ISO_TIME_RANGE")
    }
    
    v0_entry_FILE_TIME_SPAN := es[0]
    v0_entry_FILE_TIME_SPAN = utils.Trim_quoted_string(v0_entry_FILE_TIME_SPAN)
    err = utils.Iso_time_range_parser(v0_entry_FILE_TIME_SPAN)
    
    return
}
        
func (h *CefHeaderData) check_meta_DATA_TYPE() (err error) {
    return
}
        
func (h *CefHeaderData) check_meta_OBSERVATORY() (err error) {
    return
}


//- FILE_NAME="C3_CP_EDI_QZC__20111021_V01.cef"
//- 
//- START_META     = LOGICAL_FILE_ID
//-     ENTRY      = "C3_CP_EDI_QZC__20111021_V01"
//- END_META       = LOGICAL_FILE_ID
//- !
//- START_META     = VERSION_NUMBER
//-  ENTRY         = "01"
//- END_META       = VERSION_NUMBER
//- !
//- VERSION_NUMBER must be integer        
//- Version in filename /logical_file id must match VERSION_NUMBER
func (h *CefHeaderData) check_meta_VERSION_NUMBER() (err error) {
    
    v0_version_number, err := h.getMetaEntryFirstQuotedTrimed("VERSION_NUMBER")
    if err != nil {
        return
    }    

    err = utils.Integer_parser(v0_version_number)
    if err != nil {
        return
    }    

    v0_filename, err := h.getAttrFirstQuotedTrimed("FILE_NAME")
    if err != nil {
        return
    }    
    
    v0_logical, err := h.getMetaEntryFirstQuotedTrimed("LOGICAL_FILE_ID")
    if err != nil {
        return
    }
    
    //x ix := strings.LastIndex(v0_filename, ".")
    //x ix := strings.FirstIndex(v0_filename, ".")
    ix := strings.Index(v0_filename, ".")
    if ix < 0 {
        return errors.New("FILE_NAME does have .cef or .cef.gz suffix")
    }        

    fn := v0_filename[0:ix]
    if strings.HasSuffix(fn, v0_version_number) == false {
        return errors.New("FILE_NAME does have suffix (" + v0_version_number + ")")
    }
    
    if strings.HasSuffix(v0_logical, v0_version_number) == false {
        return errors.New("LOGICAL_FILE_ID does have suffix (" + v0_version_number + ")")
    }
    
    return
}
        
func (h *CefHeaderData) Meta_Checks() (err error) {
    
	mks := []string{
		"LOGICAL_FILE_ID",
		"FILE_TYPE",
		"FILE_TIME_SPAN",
		"DATA_TYPE",
		"OBSERVATORY",
		"VERSION_NUMBER",
	}

	for _, k := range mks {
		if entry, value_type, err := h.getMeta(k); err == nil {
            diag.Trace(diag.Cyan(" " + k), "Meta:ENTRY", entry)
            if value_type != nil {
                diag.Trace(diag.Cyan(" " + k), "Meta:VALUE_TYPE", value_type)
            }
		} else {
            diag.Error("Meta Checks ", err)
		}
	}
    
    err = h.check_meta_LOGICAL_FILE_ID()
    h.print_results("LOGICAL_FILE_ID checks", err) 
    
    err = h.check_meta_FILE_TYPE()
    h.print_results("FILE_TYPE checks", err) 
    
    
    err = h.check_meta_FILE_TIME_SPAN()
    h.print_results("FILE_TIME_SPAN checks", err) 
    
    h.check_meta_DATA_TYPE()
    h.check_meta_OBSERVATORY()
    
    err = h.check_meta_VERSION_NUMBER()
    h.print_results("VERSION_NUMBER checks", err) 
    
    return
}

//x h.dumpMeta()













