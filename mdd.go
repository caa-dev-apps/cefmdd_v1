package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "errors"
    "fmt"
    "io"
    "log"
    "strings"
    "strconv"
//x     "testing"
)

//- Keywords
//- Key,                    low,    high,   type,               scope,  source,     table,  comments__20160617,,
//- MISSION                 ,1,     1,      Enumerated         ,m,      CAA        ,3,      ,       ,
//- PARENT_MISSION          ,1,     1,      Enumerated         ,,       CAA        ,3,,,
//- MISSION_TIME_SPAN       ,1,     1,      ISO_TIME_RANGE     ,m,      CAA        ,3,,,

//- Enum
//- Keyword,                Definition,     Description,                                            Additional
//- MISSION,                Cluster,        ,
//- MISSION,                DoubleStar,     ,
//- PARENT_MISSION,         Cluster,        Mission or project from which the data was collected:,
//- MISSION_AGENCY,         ESA,            European Space Agency,

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

func read_csv(i_path string) chan []string {
	output := make(chan []string, 1)
    
	go func() {
		defer close(output)

		fi, err := os.Open(i_path)
		if err != nil {
            log.Fatal(err)
			return
		}
		defer fi.Close()

        r := csv.NewReader(bufio.NewReader(fi))
        ln := 0
        for {
            record, err := r.Read()
			if err == io.EOF {
				break
			}
            if err != nil {
                log.Fatal(err)
                return
            }
            
            if ln > 0 {
                for i, r := range record {
                    record[i] = strings.TrimSpace(r)
                }
                
                output <- record
            }
            
            ln++
        }
	}()

	return output
}

//x ///////////////////////////////////////////////////////////////////////////////
//x //
//x 
//x // TODO REMOVE WHEN MERGE TO REMOVE DUPLICATE
//x 
//x type KeyVal struct {
//x 	key string
//x 	val []string
//x     
//x }

///////////////////////////////////////////////////////////////////////////////
//

type KeywordType int

const (
    ENUMERATED KeywordType = iota
    FORMATTED          
    INTEGER            
    ISO_TIME           
    ISO_TIME_RANGE     
    NUMERICAL          
    STRING             
    TEXT               
)

// ------------------------------------

type KeywordData struct {
    lo              string
    hi              string
    kw_type         KeywordType
    
    scope           string
    table           string
    comment         string
    
    kw_type_parser  func(s string) (err error)
}

func (kw *KeywordData) is_range(lo, hi string) (bool) {
    return kw.lo == lo && kw.hi == hi
}

//x  &KeyVal { key: k, val: []string{ v}}
func (kw *KeywordData) test_cardinality(kv *KeyVal) (err error) {
    
    l := len(kv.val)
    //x ls := string(l)
    ls := strconv.Itoa(l)
    switch {
        case kw.is_range("0", "1") :
            if l > 1 {
                err = errors.New("Error: Keyword cardinality(0,1) Actual " + ls + " - " + kv.key)
            }
        case kw.is_range("0", "N") :
            // anything is good!
        case kw.is_range("1", "1") :
            if l != 1 {
//x                 fmt.Printf("___________________________________ 0,1 %d \n", len(kv.val))
                err = errors.New("Error: Keyword cardinality(1,1) Actual " + ls + " - " + kv.key)
            }
        case kw.is_range("1", "N") :    
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(1,N) Actual " + ls + " - " + kv.key)
            }
        case kw.is_range("0*", "1*") :
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(0*,1*) Actual " + ls + " - " + kv.key)
            }
        case kw.is_range("0*", "N*") :
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(0*,N*) Actual " + ls + " - " + kv.key)
            }
            
        default :
    }

    return 
}

///////////////////////////////////////////////////////////////////////////////
//

func (kw *KeywordData) parse_value_type(kv *KeyVal) (err error) {

    for _, v := range kv.val {
        
        v1 := strings.Trim(v, "\"'`")
        
        err = kw.kw_type_parser(v1)
        if err != nil {
            return
        }
    }

    return 
}

type EnumData struct {
    description     string
    additional      string
}

type MddData struct {
    m_keywords map[string]*KeywordData
    m_enums map[string]map[string]*EnumData
}

func NewMddData_Base(i_path_keywords, i_path_enums string) *MddData {
    
    mdd_data := &MddData{ m_keywords: make(map[string]*KeywordData), 
                          m_enums: make(map[string]map[string]*EnumData) }

    // read Enums
    for r := range read_csv(i_path_enums) {
        //x fmt.Println(r[0], r[1])
        mdd_data.add_enum(r[0], r[1], &EnumData{description: r[2], additional: r[3]})
    }

    // read Keywords
    for r := range read_csv(i_path_keywords) {
        
        l_keywordType, l_kw_type_parser, err := mdd_data.getKeywordType(r[0], r[3])
        
        if err != nil {
            log.Fatal(err)
            break
        }
        
        //x fmt.Println(r[0], r[1], r[2], r[3])
        mdd_data.add_keyword(r[0], &KeywordData{ lo : r[1], 
                                                 hi : r[2],  
                                                 kw_type : l_keywordType,    
                                                 scope : r[4],  
                                                 table : r[4],  
                                                 comment : r[5],
                                                 kw_type_parser : l_kw_type_parser,
                                                 })
    }
    
    return mdd_data
}

func (d *MddData) getKeywordType(i_kw, i_kw_type string)  (k KeywordType, f func(string) (error ), err error) {
    
    s := strings.ToLower(i_kw_type)

    switch s {
        case "enumerated" :     k = ENUMERATED;           f, err =  d.init_enum_parser(i_kw)
        case "formatted" :      k = FORMATTED;            f = formatted_parser
        case "integer" :        k = INTEGER;              f = integer_parser
        case "iso_time" :       k = ISO_TIME;             f = iso_time_parser
        case "iso_time_range" : k = ISO_TIME_RANGE;       f = iso_time_range_parser
        case "numerical" :      k = NUMERICAL;            f = numerical_parser
        case "string" :         k = STRING;               f = string_parser
        case "text" :           k = TEXT;                 f = text_parser
        default : 
            err = errors.New("Error: unknown Keyword value type: [" + i_kw_type + "]")
    }
    
    return
}

func (d *MddData) add_enum(i_keyword, 
                           i_definition string, 
                           i_enum_data *EnumData) (err error) {

    _, present := d.m_enums[i_keyword]
    if present == false {
        d.m_enums[i_keyword] = make(map[string]*EnumData)
        fmt.Println("New Enum: " + i_keyword)
    }

    
    if i_keyword == "PARAMETER_TYPE" {
        fmt.Println("#+: ", i_keyword, i_definition)
    }
    
    
    _, present = d.m_enums[i_keyword][i_definition]
    if present == true {
        return errors.New("Error: add_enum: keyword[definition] exists")
    } else {
        d.m_enums[i_keyword][i_definition] = i_enum_data
    }
    
    return
}

func (d *MddData) add_keyword(i_keyword string, 
                              i_kw_data *KeywordData) (err error) {

    _, present := d.m_keywords[i_keyword]
    if present == false {
        d.m_keywords[i_keyword] = i_kw_data
    } else {
        return errors.New("Error: add_keyword: keyword exists")
    }
    
    return
}

func (d *MddData) init_enum_parser(i_keyword string) (f func(string) (error), err error) {
    
    enums, present := d.m_enums[i_keyword]
    if present == false {
        return nil, on_enum_not_found_error(i_keyword)
    }
    
    f =  func(v string) (e error) {
        //x fmt.Printf("#?: [%s] [%s]\n", i_keyword, v)
        v1 := strings.Trim(v, "\"'`")
        
        _, p := enums[v1]
        
        //x if i_keyword == "PARAMETER_TYPE" {
        //x     fmt.Print("############ -> ")
        //x     fmt.Println(enums[v1])
        //x }
        
        if p == false {
            return on_enum_parser_error(i_keyword, v)
        }
        
        return
    }
    
    return
}
    
func (d *MddData) dump() {
    fmt.Println(d.m_keywords)
    fmt.Println(d.m_enums)
}

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

// keyword 

//- Key,                    low,    high,   type,               scope,  source,     table,  comments__20160617,,
//- MISSION                 ,1,     1,      Enumerated         ,m,      CAA        ,3,      ,       ,
// -------------------------------------------------------------------------------------------------------------------------------------------------------------
//- Enum
//- Keyword,                Definition,     Description,                                            Additional


func (d *MddData) test_input(kv *KeyVal) (err error) {

    //- k := kv.k
    //- v := kv.v [...]

    //- keyword_presence
    //- cardinality (lo, hi)
    //- kw_type
    //- check if already registered
    
    l_kw_data, present := d.m_keywords[kv.key]
    if present == false {
        return errors.New("Error: keyword not found: " + kv.key)
    } 
    
    err = l_kw_data.test_cardinality(kv)
    if err != nil {
        return 
    }
    
    err = l_kw_data.parse_value_type(kv)
    if err != nil {
        return 
    }
    
    return
}

func (d *MddData) test_input_kv1(k, 
                                 v string) (err error) {
    
    kv := &KeyVal{ key: k, val: []string{v}}
    err = d.test_input(kv)
    
    return 
}

func (d *MddData) test_input_kv2(k string, 
                                 v []string) (err error) {
    
    kv := &KeyVal{ key: k, val: v}
    err = d.test_input(kv)
    
    return 
}


func NewMddData() *MddData {
   return NewMddData_Base("C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Tables.csv", 
                          "C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Enums.csv")
}
