package readers

import (
	"fmt"
    "errors"
    "unicode"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
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

func DataRecords(i_lines chan Line, i_len int, i_data_until string) chan CefRecord {
    output := make(chan CefRecord, 16)
    l_running_len := 0

    parse_line_v2 := func(i_line string) (r_tokens []string, in_err error) {
        
        state := STX
        l_tok := ""

        for _, ch := range i_line {
            //debug 
            //d diag.Printf("%c", ch)
            //d state_0 := state

            switch state {

            case STX:                  
                switch {
                case ch == '!':
                    return
                case ch == ',':
                    in_err = errors.New(`TOO MANY COMMAS -> ` + string(ch))
                    return
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

        ix := 0
        var l_rowTokens *CefRecord = nil
        
//..    type Line struct {
//..        tag         string     // typically filename
//..        ln          int        // line number
//..        line        string     // line contents
//..    }
//..    i_lines chan readers.Line

        for l_line := range i_lines {

            if l_running_len == 0 {
                l_rowTokens = NewCefRecord(l_line)
            }

            l_tokens, err := parse_line_v2(l_line.line)
            l_rowTokens.Tokens = append(l_rowTokens.Tokens, l_tokens...)

            if err != nil {
                l_rowTokens.Err = err
                output <- *l_rowTokens
                return
            }

            l_running_len = len(l_rowTokens.Tokens)

            if l_running_len == i_len {
                output <- *l_rowTokens
                l_running_len = 0
            } else if  l_running_len > i_len {
                l_rowTokens.Err = errors.New(fmt.Sprint(`Line reader - Too many tokens - expected : `, i_len, " actual : ", l_running_len))
                output <- *l_rowTokens
                break
            }

            ix++
        }

        // check if DATA_UNTIL
        if len(i_data_until) > 0 && l_running_len == 1 && strings.Contains(i_data_until, l_rowTokens.Tokens[0]) {
            if i_data_until != l_rowTokens.Tokens[0] {
                diag.Warn("DATA_UNTIL: ", "Actual", l_rowTokens.Tokens[0], " Expected", i_data_until)
            }
        } else if l_running_len > 1 {
            l_rowTokens.Err = errors.New(fmt.Sprint(`Line reader - Too few tokens - expected : `, i_len, " actual : ", l_running_len))
            output <- *l_rowTokens
        }
        // else eof

    }()

    return output
}
