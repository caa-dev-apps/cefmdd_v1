package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////
//

//- type CefHeaderData struct {
//- 	m_state State
//- 	m_name  string
//-     m_cur* Attrs
//-
//-     m_attrs* Attrs
//-     m_meta* Meta
//-     m_vars* Vars
//- }

//- func (h *CefHeaderData) dumpAttrs()  {
//-     dumpJSON("Attrs:", h.m_attrs.m_map)
//-     fmt.Println("\n")
//- }

//- func (h *CefHeaderData) dumpMeta()  {
//-     dumpJSONMap("Meta:", h.m_meta.m_map)
//-     fmt.Println("\n")
//- }

//- func (h *CefHeaderData) dump() (err error) {
//-
//-     fmt.Println("--\n", h.m_attrs)
//-     fmt.Println("--\n", h.m_meta)
//-     fmt.Println("--\n", h.m_vars)
//-
//-
//-     fmt.Println("\n")
//-     dumpJSON("Attrs:", h.m_attrs.m_map)
//-     fmt.Println("\n")
//-     dumpJSONMap("Meta:", h.m_meta.m_map)
//-     fmt.Println("\n")
//-     dumpJSONMap("Var:", h.m_vars.m_map)
//-     fmt.Println("\n")
//-     dumpJSONList("Var:", h.m_vars.m_list)
//-     fmt.Println("\n")
//-
//- 	return err
//- }

//- func (h *CefHeaderData) getAttr(k string) (v []string, err error) {
//-
//-     v, present := h.m_attrs.m_map[k]
//-     if present == false {
//-         err = errors.New("Error: keyword not presnt " + k)
//-     }
//-
//-     return
//- }

//- func (h *CefHeaderData) getMeta(k string) (entry, value_type []string, err error) {
//-
//-     v, present := h.m_meta.m_map[k]
//-     if present == false {
//-         err = errors.New("Error: meta not presnt " + k)
//-     } else {
//-         entry, present = v.m_map[`ENTRY`]
//-         if present == false {
//-             err = errors.New("Error: meta:ENTRY not presnt " + k)
//-         } else {
//-         }
//-
//-         value_type, _ = v.m_map[`VALUE_TYPE`]
//-     }
//-
//-     return
//- }

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

func (h *CefHeaderData) checks() {
    
    //x fmt.Println("cef filename", s_args.m_filename)
    fmt.Println("cef filename", s_args.m_cefpath)
    
	aks := []string{
		"DATA_UNTIL",
		"FILE_FORMAT_VERSION",
		"FILE_NAME",
	}

	for _, k := range aks {
		if v, err := h.getAttr(k); err == nil {
			fmt.Println(k, v)
		} else {
			fmt.Println(err)
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
			fmt.Println(k, "Meta:ENTRY", entry)
            if value_type != nil {
                fmt.Println(k, "Meta:VALUE_TYPE", value_type)
            }
		} else {
			fmt.Println(err)
		}

	}
}

//x h.dumpMeta()
