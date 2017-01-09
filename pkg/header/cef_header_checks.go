package header

import (
    "fmt"
    "errors"
    "strings"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

///////////////////////////////////////////////////////////////////////////////
//

//-     mks := []string {
//-         "MISSION_KEY_PERSONNEL",
//-         "EXPERIMENT_KEY_PERSONNEL",
//-         "INSTRUMENT_NAME",
//-         "INSTRUMENT_DESCRIPTION",
//-         "DATASET_CAVEATS",
//-         "METADATA_TYPE",
//-             "LOGICAL_FILE_ID",
//-         "EXPERIMENT_DESCRIPTION",
//-         "DATASET_ID",
//-         "MIN_TIME_RESOLUTION",
//-         "PROCESSING_LEVEL",
//-         "DATASET_VERSION",
//-             "FILE_TYPE",
//-         "FILE_CAVEATS",
//-         "TIME_RESOLUTION",
//-         "DATASET_DESCRIPTION",
//-         "MISSION_REFERENCES",
//-         "MISSION_CAVEATS",
//-         "EXPERIMENT",
//-         "OBSERVATORY_DESCRIPTION",
//-         "MEASUREMENT_TYPE",
//-         "ACKNOWLEDGEMENT",
//-         "METADATA_VERSION",
//-             "FILE_TIME_SPAN",
//-         "OBSERVATORY_CAVEATS",
//-         "OBSERVATORY_REGION",
//-         "INSTRUMENT_CAVEATS",
//-         "DATASET_TITLE",
//-         "INVESTIGATOR_COORDINATES",
//-         "EXPERIMENT_CAVEATS",
//-         "MISSION_TIME_SPAN",
//-         "MISSION_AGENCY",
//-         "MISSION_REGION",
//-         "EXPERIMENT_REFERENCES",
//-         "INSTRUMENT_TYPE",
//-             "DATA_TYPE",
//-         "MAX_TIME_RESOLUTION",
//-         "MISSION",
//-         "GENERATION_DATE",
//-         "MISSION_DESCRIPTION",
//-             "OBSERVATORY",
//-         "OBSERVATORY_TIME_SPAN",
//-         "CONTACT_COORDINATES",
//-             "VERSION_NUMBER",
//-     }

///////////////////////////////////////////////////////////////////////////////
//

func (h *CefHeaderData) getAttrFirstQuoted(key string) (v0 string, err error) {

    vs, err := h.getAttr(key);
    if err == nil {  
        v0 = vs[0]
        if utils.Is_quoted_string(v0) == false {
            err = errors.New("error: " + key + "attribute is unquoted string ")
        } 
    }
        
    return
}


func (h *CefHeaderData) getMetaEntryFirstQuoted(key string) (v0 string, err error) {

    vs, err := h.getMetaEntry(key);
    if err == nil {  
        v0 = vs[0]
        if utils.Is_quoted_string(v0) == false {
            err = errors.New("error: " + key + "META ENTRY is unquoted string ")
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
        err = errors.New("error: Attr: " + key + " length = 0")
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
        err = errors.New("error: Meta:ENTRY: " + key + " length = 0")
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
    //x l_args := utils.GetCefArgs()
    l_filename := utils.GetCefArgs().GetFilename()
//x     if l_args.m_filename != v0 {
//x         err = errors.New("error: FILE_NAME attribute mismatches - actual (" + l_args.m_filename + ")")
//x     } 

    if l_filename != v0 {
        err = errors.New("error: FILE_NAME attribute(" + v0 + ") - mismatches actual (" + l_filename + ")")
    } 
    
    return
}

func (h *CefHeaderData) check_attr_FILE_FORMAT_VERSION() (err error) {
    
    fmt.Println(diag.BoldMagenta("TODO check_attr_FILE_FORMAT_VERSION"))  
    
    return
}

func (h *CefHeaderData) check_attr_DATA_UNTIL() {
    
    fmt.Println(diag.BoldMagenta("TODO check_attr_DATA_UNTIL"))  
    return
}

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
        err = errors.New("error: DATASET_ID length = 0")
        return
    } 
    
    v0_logical = utils.Trim_quoted_string(v0_logical)
    if strings.HasPrefix(v0_logical, v0_dataset) == false {
        err = errors.New("error: DATASET_ID (" + v0_dataset + ") is not a prefix of LOGICAL_FILE_ID (" + v0_logical + ")" )
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
        return errors.New("error: FILE_NAME length = 0")
    } 
    
    v0_file_type = utils.Trim_quoted_string(v0_file_type)
    if len(v0_file_type) == 0 {
        return errors.New("error: FILE_TYPE length = 0")
    } 
    
    v0_logical = utils.Trim_quoted_string(v0_logical)
    if len(v0_logical) == 0 {
        return errors.New("error: LOGICAL_FILE_ID length = 0")
    } 
    
    
    //x ix := strings.FirstIndex(v0_filename, ".")
    ix := strings.Index(v0_filename, ".")
    if ix < 0 {
        return errors.New("error: FILE_NAME does have .cef or .cef.gz suffix")
    }        

    ext0 := v0_filename[ix+1:]
    if strings.HasPrefix(ext0, v0_file_type) == false {
        return errors.New("error: FILE_TYPE does not follow first '.' in actual filename" + "||" + ext0 + "||" + v0_file_type)
    }
    
    fn := v0_logical + `.` + v0_file_type
    // allow for .gz
    if strings.HasSuffix(v0_filename, fn) == false {
        fmt.Println(v0_filename)
        fmt.Println(fn)
        
        return errors.New("error: FILE_NAME does not match LOGICAL_FILE_ID + FILE_TYPE  || " + v0_filename  + " || " + fn)
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
        return errors.New("error: META-ENTRY for FILE_TIME_SPAN Missing")
    }
    
    if len(vs) == 0 {
        return errors.New("error: META-VALUE_TYPE for FILE_TIME_SPAN Missing")
    }
    
    v1_value_type_FILE_TIME_SPAN := vs[0]
    if "ISO_TIME_RANGE" != v1_value_type_FILE_TIME_SPAN {
        return errors.New("error: META-VALUE_TYPE for FILE_TIME_SPAN not equal to ISO_TIME_RANGE")
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
        return errors.New("error: FILE_NAME does have .cef or .cef.gz suffix")
    }        

    fn := v0_filename[0:ix]
    if strings.HasSuffix(fn, v0_version_number) == false {
        return errors.New("error: FILE_NAME does have suffix (" + v0_version_number + ")")
    }
    
    if strings.HasSuffix(v0_logical, v0_version_number) == false {
        return errors.New("error: LOGICAL_FILE_ID does have suffix (" + v0_version_number + ")")
    }
    
    return
}
        
func (h *CefHeaderData) print_results(about string, err error) {

//x     if err != nil {
//x         fmt.Println(diag.BoldRed("error-" + about), err)        
//x     } else {
//x         fmt.Println(diag.BoldGreen("ok-" + about))        
//x     }
    if err != nil {
        diag.Error(diag.BoldRed("error:"), about, err)   
    } else {
        diag.Info(diag.BoldGreen("ok:"), about)  
    }
}        
        
        
func (h *CefHeaderData) Checks() (err error) {
    
    //x fmt.Println("cef filename", s_args.m_filename)
    //xfmt.Println("cef cefpath", *s_args.m_cefpath)
    //x fmt.Println("cef cefpath2", s_args.m_cefpath)
    
	aks := []string{
		"FILE_NAME",
		"FILE_FORMAT_VERSION",
		"DATA_UNTIL",
	}

	for _, k := range aks {
		if v, err := h.getAttr(k); err == nil {
            diag.Info(diag.Cyan(k), v)
		} else {
            diag.Error(diag.BoldRed(err))
		}
	}

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
            diag.Info(diag.Blue(k), "Meta:ENTRY", entry)
            if value_type != nil {
                diag.Info(diag.Yellow(k), "Meta:VALUE_TYPE", value_type)
            }
		} else {
            diag.Error(diag.BoldRed(err))
		}
	}
    
    err = h.check_attr_FILE_NAME()
    h.print_results("FILE_NAME checks", err) 
    
    h.check_attr_FILE_FORMAT_VERSION()
    h.check_attr_DATA_UNTIL()

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













