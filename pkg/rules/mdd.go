package rules

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
    "github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
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

func (kw *KeywordData) test_cardinality(kv *readers.KeyVal) (err error) {

    l := len(kv.Val)
    ls := strconv.Itoa(l)
    switch {
        case kw.is_range("0", "1") :
            if l > 1 {
                err = errors.New("Error: Keyword cardinality(0,1) Actual " + ls + " - " + kv.Key + fmt.Sprintf(" %s ",kv.Val))
            }
        case kw.is_range("0", "N") :
            // anything is good!
        case kw.is_range("1", "1") :
            if l != 1 {
                err = errors.New("Error: Keyword cardinality(1,1) Actual " + ls + " - " + kv.Key + fmt.Sprintf(" %s ",kv.Val))
            }
        case kw.is_range("1", "N") :    
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(1,N) Actual " + ls + " - " + kv.Key)
            }
        case kw.is_range("0*", "1*") :
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(0*,1*) Actual " + ls + " - " + kv.Key)
            }
        case kw.is_range("0*", "N*") :
            if l == 0 {
                err = errors.New("Error: Keyword cardinality(0*,N*) Actual " + ls + " - " + kv.Key)
            }
            
        default :
    }
//20191111 smcc added 
        if err != nil {
            log.Fatal(err)
        }
//20191111 smcc added 

    return 
}

///////////////////////////////////////////////////////////////////////////////
//

func (kw *KeywordData) parse_value_type(kv *readers.KeyVal) (err error) {

    for _, v := range kv.Val {
        
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
        mdd_data.add_enum(r[0], r[1], &EnumData{description: r[2], additional: r[3]})
    }

    // read Keywords
    for r := range read_csv(i_path_keywords) {
        l_keywordType, l_kw_type_parser, err := mdd_data.getKeywordType(r[0], r[3])
        if err != nil {
            log.Fatal(err)
            break
        }

//SMCC QWE HERE IS WHERE THE CSV FOR CARD IS SETUP        

//fmt.Println("ADDING" , r[0] , r[1], r[2])
        
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
        case "enumerated" : {
    k = ENUMERATED;           
  f, err =  d.init_enum_parser(i_kw)
}
        case "formatted" :      k = FORMATTED;            f = utils.Formatted_parser
        case "integer" :        k = INTEGER;              f = utils.Integer_parser
        case "iso_time" :       k = ISO_TIME;             f = utils.Iso_time_parser
        case "iso_time_range" : k = ISO_TIME_RANGE;       f = utils.Iso_time_range_parser
        case "numerical" :      k = NUMERICAL;            f = utils.Numerical_parser
        case "string" :         k = STRING;               f = utils.String_parser
        case "text" :           k = TEXT;                 f = utils.Text_parser
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
        return nil, utils.On_enum_not_found_error(i_keyword)
    }
    
    f = func(v string) (e error) {
        v1 := strings.Trim(v, "\"'`")

        if i_keyword == "COORDINATE_SYSTEM" {
	if (strings.Contains(v1,">")){
	            vs := strings.Split(v1, ">")
	            _, p := enums[vs[0]]
		 	if (p == false) {
			      return utils.On_enum_parser_error(i_keyword, v)
			     }	
		            if (len(vs[1]) < 1) {
			        err = errors.New(fmt.Sprintf("Error: Incomplete value %s for %s ",vs, i_keyword))
				return err	
		            }
	}  else {	
	//This COORD SYSTEM HAS NO DESCRIPTOR ATTACHED
            _, p := enums[v]
		 if (p == false) {
	      return utils.On_enum_parser_error(i_keyword, v)
	     }	

	}
        } else if i_keyword == "DATA_TYPE" {
            vs := strings.Split(v1, ">")
            _, p := enums[vs[0]]

            if p == false {
                //x 
		return utils.On_enum_parser_error(i_keyword, v)
                //20181009 removed diag.Warn("Parser, Enum description mismatch ", i_keyword, v)
            }

            //x if (len(vs) > 1 && es.description != vs[1]) {
            //x     return utils.On_enum_description_error(i_keyword, v)
            //x } 
        } else if i_keyword == "TARGET_SYSTEM" {  //20181009 SNCC Added 
            vs := strings.Split(v1, ">")
            _, p := enums[vs[0]]

            if p == false {
                return utils.On_enum_parser_error(i_keyword, v)
//                diag.Warn("Parser, Enum description mismatch ", i_keyword, v)
            }

            //x if (len(vs) > 1 && es.description != vs[1]) {
            //x     return utils.On_enum_description_error(i_keyword, v)
            //x } 
        } else if i_keyword == "SI_CONVERSION" {
	   v1_orig:=v1
	   v1 = RemoveCommentInBrackets(v1)
//	   fmt.Println("SI:",v1)

            vs := strings.Split(v1, ">")
            l := len(vs) 

            if l==0 || l > 2 {
                //2017-02-11 return utils.On_enum_parser_error(i_keyword, v)
                err = errors.New("Error: SI_CONVERSION format error ["  + v1_orig + "]")
		log.Fatal(err)
            }

	   if !(isNumeric(vs[0])){
                err = errors.New("Error: SI_CONVERSION format error, must have numerical value [" + v1_orig + "]")
	        log.Fatal(err)
  	   }

	if (l==2){
	if (len(vs[1]) == 0){
	// There is no unit given
	  err = errors.New("Error: SI_CONVERSION invalid, no unit given [" + v1_orig + "]")
	  log.Fatal(err)
	
	}else {
		arrUnits := strings.Split(strings.TrimSpace(vs[1])," ")  // This should be each SI unit

		for _, element := range arrUnits {
			a := strings.FieldsFunc(element,Split)
			if (len(a) <=2 ){
//				fmt.Println("Correct number of values in ", element)			
			} else{
		                err = errors.New("Error: SI_CONVERSION invalid, incorrect number of values [" + v1_orig + "]")
			        log.Fatal(err)
			}
//	fmt.Println("A:",len(a))
			//Now validate unit

	            	 _, p := enums[a[0]]
	            	if p == false {
	                	err= utils.On_enum_parser_error(i_keyword, a[0])
			        log.Fatal(err)
		            }
			//if it is in the format 1>x^123
			if (len(a) ==2){		
			   if !(isNumeric(a[1])){
		                err = errors.New("Error:  SI_CONVERSION must have numerical value [" + v1_orig + "]")
			        log.Fatal(err)
		  	   }
			} 
	
		}
	}
}
//            ix := 0
//            if l == 2 {
//                ix = 1
//            }

//2019 smcc			 fmt.Println("V:[", v, "]","IX:[", vs[ix])

        } else if strings.EqualFold(v, "MULTIPLE") && (i_keyword == "EXPERIMENT" || i_keyword == "INSTRUMENT_NAME" || i_keyword == "OBSERVATORY") {
            // pass
        } else {
            _, p := enums[v1]

            if p == false {
	            log.Fatal(utils.On_enum_parser_error(i_keyword, v))

            } 
        }

        return
    }

    
    return
}
    
func (d *MddData) dump() {
    diag.Trace(fmt.Sprintf("%#v", d.m_keywords))
    diag.Trace(fmt.Sprintf("%#v", d.m_enums))
}


func (d *MddData) BuildMDDCheck(strType string) (arrReturn []string) {

	for k, v := range d.m_keywords {
		if ((v.scope == strType) && (v.lo == "1")){
			arrReturn = append(arrReturn,k)
		}
//		fmt.Println(fmt.Sprintf("K %s MIN:%s MAX:%s SCOPE:%s",k,v.lo,v.hi,v.scope))
	}
//	fmt.Println("******************BEGIN***************")
//	fmt.Println("******************END***************")
return arrReturn
}

//    fmt.Println(fmt.Sprintf("%#v", d.m_keywords))
//    fmt.Println(fmt.Sprintf("%#v", d.m_enums))

// keyword 

//- Key,                    low,    high,   type,               scope,  source,     table,  comments__20160617,,
//- MISSION                 ,1,     1,      Enumerated         ,m,      CAA        ,3,      ,       ,
// -------------------------------------------------------------------------------------------------------------------------------------------------------------
//- Enum
//- Keyword,                Definition,     Description,                                            Additional


func (d *MddData) Test_input(kv *readers.KeyVal) (err error) {

    //- k := kv.k
    //- v := kv.v [...]

    //- keyword_presence
    //- cardinality (lo, hi)
    //- kw_type
    //- check if already registered

//2018    fmt.Println("KEYV: " , kv.Key)

    l_kw_data, present := d.m_keywords[kv.Key]


//fmt.Println("DATA", d.m_keywords[kv.Key])

    if present == false {
        return errors.New("Error: keyword not found: " + kv.Key)
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

func (d *MddData) Test_input_kv1(k, 
                                 v string) (err error) {
    kv := &readers.KeyVal{ Key: k, Val: []string{v}}
    err = d.Test_input(kv)
    
    return 
}

func (d *MddData) Test_input_kv2(k string, 
                                 v []string) (err error) {
    kv := &readers.KeyVal{ Key: k, Val: v}
    err = d.Test_input(kv)
    
    return 
}

func NewMddData() *MddData {
   
    return NewMddData_Base(utils.CefMddDir() + `/Keywords.csv`, 
                           utils.CefMddDir() + `/Enums.csv`)
}

func RemoveCommentInBrackets(inStr string)(string){
	for (strings.Contains(inStr,"(")) {
		inStr = string(inStr[0:strings.Index(inStr,"(")]) + string(inStr[strings.Index(inStr,")")+1:len(inStr)])
	}
	return inStr
}

func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func Split(r rune) bool {
    return r =='^' 
}
