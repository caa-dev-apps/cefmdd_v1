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

///////////////////////////////////////////////////////////////////////////////
//

// TODO REMOVE WHEN MERGE TO REMOVE DUPLICATE

type KeyVal struct {
	key string
	val []string
    
}

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

func getKeywordType(s string)  (k KeywordType, err error) {
    
    s = strings.ToLower(s)

    switch s {
        case "enumerated" :     k = ENUMERATED
        case "formatted" :      k = FORMATTED
        case "integer" :        k = INTEGER
        case "iso_time" :       k = ISO_TIME
        case "iso_time_range" : k = ISO_TIME_RANGE
        case "numerical" :      k = NUMERICAL
        case "string" :         k = STRING
        case "text" :           k = TEXT    
        default : 
            err = errors.New("Error: unknown Keyword value type: [" + s + "]")
    }
    
    return
}

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
    ls := string(l)
    switch {
        case kw.is_range("0", "1") :
            if l > 1 {
                err = errors.New("Error: Keyword cardinality(0,1) Actual " + ls + " - " + kv.key)
            }
        case kw.is_range("0", "N") :
            // anything is good!
        case kw.is_range("1", "1") :
            if l != 1 {
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
        
        fmt.Printf("-----> %s %p", v, kw.kw_type_parser)
        
        err = kw.kw_type_parser(v)
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
    m_keywords map[string]KeywordData
    m_enums map[string]map[string]EnumData
}

func NewMddData_Base(i_path_keywords, i_path_enums string) *MddData {
    
    wd, err := os.Getwd()
    if err != nil {
        return nil
    }
    
    fmt.Println("_________________________________")
    fmt.Println("Current Working Directory = " + wd)
    
    
    mdd_data := &MddData{ m_keywords: make(map[string]KeywordData), 
                          m_enums: make(map[string]map[string]EnumData) }
    
    //x for r := range read_csv("C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Tables.csv") {
    for r := range read_csv(i_path_keywords) {
        
        l_keywordType, err := getKeywordType(r[3])
        
        if err != nil {
            log.Fatal(err)
            break
        }
        
        //x fmt.Println(r[0], r[1], r[2], r[3])
        mdd_data.add_keyword(r[0], &KeywordData{ lo : r[1], hi : r[2],  kw_type : l_keywordType,    scope : r[4],  table : r[4],  comment : r[5] })
    }
    
    //x fmt.Println("\n________________________________________\n")
    //x for r := range read_csv("C:\\work.dev.go\\src\\github.com\\caa-dev-apps\\cefmdd_v1\\_mdd-csv\\mdd__20160617.xlsx - Enums.csv") {
    for r := range read_csv(i_path_enums) {
        //x fmt.Println(r[0], r[1])
        mdd_data.add_enum(r[0], r[1], &EnumData{description: r[2], additional: r[2]})
    }

    // associate parser with each Keyword - 
    mdd_data.init_parsers()
    
    return mdd_data
    //x mdd_data.dump()
}


func (d *MddData) add_enum(i_keyword, 
                           i_definition string, 
                           i_enum_data *EnumData) (err error) {

    _, present := d.m_enums[i_keyword]
    if present == false {
        d.m_enums[i_keyword] = make(map[string]EnumData)
    }

    _, present = d.m_enums[i_keyword][i_definition]
    if present == true {
        return errors.New("Error: add_enum: keyword[definition] exists")
    } else {
        d.m_enums[i_keyword][i_definition] = *i_enum_data
    }
    
    return
}

func (d *MddData) add_keyword(i_keyword string, 
                              i_kw_data *KeywordData) (err error) {

    _, present := d.m_keywords[i_keyword]
    if present == false {
        d.m_keywords[i_keyword] = *i_kw_data
    } else {
        return errors.New("Error: add_keyword: keyword exists")
    }
    
    return
}

//x     m_keywords map[string]KeywordData
//x     m_enums map[string]map[string]EnumData
//x 
//x     kw_type

func (d *MddData) init_enum_parser(i_keyword string) (f func(string) (error), err error) {
    
    enums, present := d.m_enums[i_keyword]
    if present == false {
        return nil, errors.New("Enum not found for KeyWord -> " + i_keyword)
    }
    
    f =  func(v string) (e error) {
        _, p := enums[v]
        if p == false {
            return errors.New("Enum value not found for KeyWord -> " + i_keyword + " value: " + v)
        }
        
        return
    }
    
    return
}
    
//x func (d *MddData) init_parsers() (err error) {
//x 
//x     for kw, v := range d.m_keywords {
//x 
//x         switch v.kw_type {
//x             case ENUMERATED :
//x                 if v.kw_type_parser, err =  d.init_enum_parser(kw); err != nil {
//x                     return
//x                 }
//x             case FORMATTED :
//x                 v.kw_type_parser =  formatted_parser
//x             case INTEGER :
//x                 v.kw_type_parser =  integer_parser
//x             case ISO_TIME :
//x                 v.kw_type_parser =  iso_time_parser
//x             case ISO_TIME_RANGE :
//x                 v.kw_type_parser =  iso_time_range_parser
//x             case NUMERICAL :
//x                 v.kw_type_parser =  numerical_parser
//x             case STRING :
//x                 v.kw_type_parser = string_parser
//x             case TEXT :
//x                 v.kw_type_parser = text_parser
//x             default :
//x                 err = errors.New("Error: unknown parser type")
//x         }
//x     }
//x 
//x     return
//x }

func (d *MddData) init_parsers() (err error) {

    for kw, _ := range d.m_keywords {

//x         d.m_keywords[kw]
        switch d.m_keywords[kw].kw_type {
            case ENUMERATED :
                if d.m_keywords[kw].kw_type_parser, err =  d.init_enum_parser(kw); err != nil {
                    return
                }
            case FORMATTED :
                d.m_keywords[kw].kw_type_parser =  formatted_parser
            case INTEGER :
                d.m_keywords[kw].kw_type_parser =  integer_parser
            case ISO_TIME :
                d.m_keywords[kw].kw_type_parser =  iso_time_parser
            case ISO_TIME_RANGE :
                d.m_keywords[kw].kw_type_parser =  iso_time_range_parser
            case NUMERICAL :
                d.m_keywords[kw].kw_type_parser =  numerical_parser
            case STRING :
                d.m_keywords[kw].kw_type_parser = string_parser
            case TEXT :
                d.m_keywords[kw].kw_type_parser = text_parser
            default :
                err = errors.New("Error: unknown parser type")
        }
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


