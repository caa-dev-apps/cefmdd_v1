package readers

import (
	"fmt"
    "errors"
    "unicode"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/utils"
    "strings"
)

type ROWState int

const (
    STX ROWState = iota
    TOK
    TOK_2
    QSTR
    QSTR_2
    ERROR
)

// TOK        TOK    TOK  TOK     QTOK    TOK
// DATE-TIME, FLOAT, INT, DOUBLE, STRING, EOL_MARKER

func (s ROWState)state_str() (str string) {

    switch s {
    case STX:                   return "STX        "
    case TOK:                   return "TOK         "
    case TOK_2:                 return "TOK_2       "
    case QSTR:                  return "QSTR        "
    case QSTR_2:                return "QSTR_2      "
    case ERROR:                 return "ERROR       "
    default:                    return "????????    "
    }

    return
}

func (s ROWState)state_diag(ch rune, s2 ROWState, tok string) {
    diag.Println(string(ch), s.state_str(), s2.state_str(), tok)
}

//smcc 20210531
func isASCII(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] > unicode.MaxASCII {
            return false
        }
    }
    return true
}
//smcc


type CefRecord struct {
    Err error
    Tokens []string
    //- Line string       
    Line Line       
}

func (row CefRecord) String() string {
     return fmt.Sprintf("In: %#v\nTokens: %#v\n", row.Line, row.Tokens)
}

func (row *CefRecord) Append(t string)  {
     row.Tokens = append(row.Tokens, t)
}

func NewCefRecord(l Line) (*CefRecord) {
    return &CefRecord{  
        Tokens: make([]string, 0),
        Line : l,
    }
}

func DataRecords(i_lines chan Line, i_len int, i_data_until, i_end_of_line_marker string) chan CefRecord {
    //x diag.Info("LINE-LENGTH = ", i_len)
// fmt.Println("LINE-LENGTH = ", i_len)

    i_data_until = strings.TrimSpace(i_data_until)
    output := make(chan CefRecord, 16)
    l_running_len := 0
    l_check_eolm := len(i_end_of_line_marker) == 1


    parse_line_v2 := func(i_line string) (r_tokens []string, in_err error) {

        state := STX
        l_tok := ""
   	intDataLines := 0
        for _, ch := range i_line {
            //debug 

//             fmt.Println("HERE:", i_line)
//             fmt.Println("%c", ch)
//             state_0 := state

    intDataLines ++
            switch state {

            case STX:                  
                switch {
                case ch == '!':
                    return
                case ch == ',':
                    if(l_running_len == 0) {
                        in_err = errors.New(`TOO MANY COMMAS -> ` + string(ch))
                        return
                    }
                case unicode.IsSpace(ch):   // tab, space \n \r
                    // No Change
                case ch == '"':
                    l_tok = string(ch)
                    state = QSTR
                default:
                    l_tok = string(ch)
                    state = TOK
                }
            case TOK:                   
                switch {
                case ch == '!':
                    r_tokens = append(r_tokens, l_tok)
                    return
                case ch == ',':
                    r_tokens = append(r_tokens, l_tok)
                    l_tok = ""
                    state = STX
//0             case ch == '$':
                case l_check_eolm && ch == rune(i_end_of_line_marker[0]):
                    // possible 2 more cells to add .. hence -2
                    if (len(r_tokens) + l_running_len) == (i_len - 2) {
                        r_tokens = append(r_tokens, l_tok)
                        l_tok = string(ch)
                        state = TOK
                    } else {
                        in_err = errors.New(`LINE READER TOK, EOL MARKER out of position, ` + string(ch))
                        return
                    }

                case unicode.IsSpace(ch):   // tab, space \n \r
                    r_tokens = append(r_tokens, l_tok)
                    l_tok = ""
                    state = TOK_2
                default:
                    l_tok += string(ch)
                }

            case TOK_2:                   
                switch {
                case ch == '!':
                    return
                case ch == ',':
                    state = STX
                case unicode.IsSpace(ch):   // tab, space \n \r
                    // noop
                default:
                    // allow for end of record marker without previous comma
                    // e.g. ... " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -10, -1000000000 $
                    if (len(r_tokens) + l_running_len) == (i_len -1) {
                        l_tok = string(ch)
                        state = TOK
                        if ch == '"' {
                            state = QSTR
                        } 
                    } else {
                        in_err = errors.New(`LINE READER TOK_2 (expected ',',space,tab,!) - got ` + string(ch))
                        return
                    }
                }

            case QSTR:     
                switch {
                case ch == '"':
                    l_tok += string(ch)
                    r_tokens = append(r_tokens, l_tok)
                    state = TOK_2
                default:
                    l_tok += string(ch)
                }

            default:                    
            }

            //debug 
            //d state_diag(ch , state_0, state, l_tok)
            //d state_0.state_diag(ch, state, l_tok)
        }

        if state == TOK {
            r_tokens = append(r_tokens, l_tok)
        }

        return
    }
    go func() {
        defer close(output)

	intEmptyLine:=0
        ix := 0

        var l_rowTokens *CefRecord = nil
        
//..    type Line struct {
//..        tag         string     // typically filename
//..        ln          int        // line number
//..        line        string     // line contents
//..    }
//..    i_lines chan readers.Line

        for l_line := range i_lines {
	    intEmptyLine=0

            if l_running_len == 0 {
                l_rowTokens = NewCefRecord(l_line)
            }

            l_tokens, err := parse_line_v2(l_line.line)
	    l_line_orig := l_line.line
//20210610 smcc added
clean := strings.Map(func(r rune) rune {
    if unicode.IsGraphic(r) {
        return r
    }
    return -1
}, l_line.line)
if (len(l_line_orig) - len(clean) > 2){

//fmt.Println("LEN ORIG",len(l_line_orig))
//fmt.Println("LEN AFTER",len(clean))
//fmt.Println("LINE ORIG:",l_line_orig)
//fmt.Println("LINE AFTER:",clean)

            l_rowTokens.Err = errors.New(fmt.Sprint(`Line reader - unexpected characters detected in line`,l_line_orig))
            output <- *l_rowTokens
	return 
}

            l_line.line = strings.TrimSpace(l_line.line)
//fmt.Println("Updated: ", l_line.line)
            l_rowTokens.Tokens = append(l_rowTokens.Tokens, l_tokens...)

            if err != nil {
                l_rowTokens.Err = err
                output <- *l_rowTokens
                return
            }
            l_running_len = len(l_rowTokens.Tokens)

            if l_running_len == i_len {
                 if (strings.Contains(l_line_orig,"\n") == true){
	                output <- *l_rowTokens
	                l_running_len = 0
		}else{
            	  	l_rowTokens.Err = errors.New(fmt.Sprint("Missing EOR ", l_line.line))
	            	output <- *l_rowTokens
			return
		}
            } else if  l_running_len > i_len {
                l_rowTokens.Err = errors.New(fmt.Sprint(`Line reader - Too many tokens - expected : `, i_len, " actual : ", l_running_len))
                output <- *l_rowTokens
                break
            } else if l_running_len == 0 {
		if (len(i_data_until) > 2){
		 if (len(strings.TrimSpace(l_line.line))==0){
		    intEmptyLine = 1
		} else  if ((len(strings.TrimSpace(l_line.line))==0) &&  (l_line.line[0]!= 33)){ //it is a space
		    intEmptyLine = 1
	}

		}

	    }

            ix++

//	if (len(i_data_until) > 2) && strings.Contains(utils.Trim_quoted_string(l_line.line),utils.Trim_quoted_string(i_data_until)) {
//		if (utils.Trim_quoted_string(l_line.line) == utils.Trim_quoted_string(i_data_until)){
//		intDataUntil_LineNo = l_line.ln
//
//		if (i_data_until != l_rowTokens.Tokens[0]){
//		          diag.Warn("DATA_UNTIL: ", "Actual", l_rowTokens.Tokens[0], " Expected : ", i_data_until)
//	}
//		}
//	}

        }

        // check if DATA_UNTIL

        if len(i_data_until) > 0 && l_running_len == 1{
		if (strings.TrimSpace(utils.Trim_quoted_string(l_rowTokens.Tokens[0])) != utils.Trim_quoted_string(i_data_until)) {
	            l_rowTokens.Err = errors.New(fmt.Sprint("DATA_UNTIL: ", "Actual ", l_rowTokens.Tokens[0], " Expected ", utils.Trim_quoted_string(i_data_until)))
	            output <- *l_rowTokens
            }  else if (strings.TrimSpace(utils.Trim_quoted_string(l_rowTokens.Tokens[0])) != strings.TrimSpace(l_rowTokens.Tokens[0])) {
	            l_rowTokens.Err = errors.New(fmt.Sprint("DATA_UNTIL: ", "Actual ", l_rowTokens.Tokens[0], " Expected ", utils.Trim_quoted_string(l_rowTokens.Tokens[0])))
	            output <- *l_rowTokens
		}  else if (utils.Trim_quoted_string(i_data_until) == i_data_until) {
	            l_rowTokens.Err = errors.New(fmt.Sprint("DATA_UNTIL: ", i_data_until, " should be in quotes"))
	            output <- *l_rowTokens
		}

        } else if l_running_len > 0 {  // was 1
            l_rowTokens.Err = errors.New(fmt.Sprint(`Line reader - Too few tokens - expected : `, i_len, " actual : ", l_running_len))
            output <- *l_rowTokens

        }
        // else eof
	if (intEmptyLine==1){
	            l_rowTokens.Err = errors.New(`DATA_UNTIL0 value undetected, expected value: ` + utils.Trim_quoted_string(i_data_until))
		    output <- *l_rowTokens}

    }()

    return output
}
