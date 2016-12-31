package readers

import (
	"fmt"
    "errors"
    "unicode"
)

type ROWState int

const (
    STX ROWState = iota
    TOK
    TOK_2
    QSTR
    QSTR_2
    //x COMMENT
    ERROR
    //x EOL
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

    fmt.Println(string(ch), s.state_str(), s2.state_str(), tok)
}


type RowTokens struct {
    Err error
    Tokens []string
    //x Line string       // 
    Line Line       // 
}

func (row RowTokens) String() string {
     return fmt.Sprintf("In: %#v\nTokens: %#v\n", row.Line, row.Tokens)
}

func (row *RowTokens) Append(t string)  {
     row.Tokens = append(row.Tokens, t)
}

//x func NewRowTokens0() (*RowTokens) {
//x     return &RowTokens{
//x         Tokens: make([]string, 0),
//x     }
//x }

func NewRowTokens(l Line) (*RowTokens) {
    return &RowTokens{  
        Tokens: make([]string, 0),
        Line : l,
    }
}


//x func RowTokeniser(i_lines chan string, i_len int) chan RowTokens {
//x     output := make(chan RowTokens, 16)
//x     l_running_len := 0
//x 
//x     parse_line_v2 := func(i_line string) (r_tokens []string, in_err error) {
//x         
//x         state := STX
//x         l_tok := ""
//x 
//x         for _, ch := range i_line {
//x             //debug 
//x             //d fmt.Printf("%c", ch)
//x             //d state_0 := state
//x 
//x             switch state {
//x 
//x             case STX:                  
//x                 switch {
//x                 case ch == '!':
//x                     return
//x                 case ch == ',':
//x                     in_err = errors.New(`TOO MANY COMMAS -> ` + string(ch))
//x                     return
//x                 case unicode.IsSpace(ch):   // tab, space \n \r
//x                     // No Change
//x                 case ch == '"':
//x                     l_tok = string(ch)
//x                     state = QSTR
//x                 default:
//x                     l_tok = string(ch)
//x                     state = TOK
//x                 }
//x 
//x             case TOK:                   
//x                 switch {
//x                 case ch == '!':
//x                     r_tokens = append(r_tokens, l_tok)
//x                     return
//x                 case ch == ',':
//x                     r_tokens = append(r_tokens, l_tok)
//x                     l_tok = ""
//x                     state = STX
//x                 case unicode.IsSpace(ch):   // tab, space \n \r
//x                     r_tokens = append(r_tokens, l_tok)
//x                     l_tok = ""
//x                     state = TOK_2
//x                 default:
//x                     l_tok += string(ch)
//x                 }
//x 
//x             case TOK_2:                   
//x                 switch {
//x                 case ch == '!':
//x                     return
//x                 case ch == ',':
//x                     state = STX
//x                 case unicode.IsSpace(ch):   // tab, space \n \r
//x                     // noop
//x                 default:
//x                     // allow for end of record marker without previous comma
//x                     // e.g. ... " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -10, -1000000000 $
//x                     if (len(r_tokens) + l_running_len) == (i_len -1) {
//x                         l_tok = string(ch)
//x                         state = TOK
//x                         if ch == '"' {
//x                             state = QSTR
//x                         } 
//x                     } else {
//x                         in_err = errors.New(`LINE READER TOK_2 (expected ',',space,tab,!) - got ` + string(ch))
//x                         return
//x                     }
//x                 }
//x 
//x             case QSTR:     
//x                 switch {
//x                 case ch == '"':
//x                     l_tok += string(ch)
//x                     r_tokens = append(r_tokens, l_tok)
//x                     state = TOK_2
//x                 default:
//x                     l_tok += string(ch)
//x                 }
//x 
//x             default:                    
//x             }
//x 
//x             //debug 
//x             //d state_diag(ch , state_0, state, l_tok)
//x             //d state_0.state_diag(ch, state, l_tok)
//x         }
//x 
//x         if state == TOK {
//x             r_tokens = append(r_tokens, l_tok)
//x         }
//x 
//x         return
//x     }
//x 
//x     go func() {
//x         defer close(output)
//x 
//x         ix := 0
//x         l_running_len = 0
//x         l_rowTokens := NewRowTokens()
//x         
//x         for l_line := range i_lines {
//x 
//x             l_tokens, err := parse_line_v2(l_line)
//x             l_rowTokens.Tokens = append(l_rowTokens.Tokens, l_tokens...)
//x 
//x             if err != nil {
//x                 l_rowTokens.Err = err
//x                 output <- *l_rowTokens
//x                 return
//x             }
//x 
//x             l_running_len = len(l_rowTokens.Tokens)
//x 
//x             if l_running_len == i_len {
//x                 output <- *l_rowTokens
//x                 l_running_len = 0
//x                 l_rowTokens = NewRowTokens()
//x             }
//x 
//x             if l_running_len > i_len {
//x                 l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too many tokens - expected : `, i_len, " actual : ", l_running_len))
//x                 output <- *l_rowTokens
//x                 break
//x             }
//x 
//x             ix++
//x         }
//x 
//x         if l_running_len > 0 && l_running_len < i_len {
//x             l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too few tokens - expected : `, i_len, " actual : ", l_running_len))
//x             output <- *l_rowTokens
//x         }
//x 
//x     }()
//x 
//x     return output
//x }

//..    //x type Line struct {
//..    //x     tag         string     // typically filename
//..    //x     ln          int        // line number
//..    //x     line        string     // line contents
//..    //x }
//..    i_lines chan readers.Line

//x func RowTokeniser(i_lines chan string, i_len int) chan RowTokens {
func RowTokeniser(i_lines chan Line, i_len int) chan RowTokens {
    output := make(chan RowTokens, 16)
    l_running_len := 0

    parse_line_v2 := func(i_line string) (r_tokens []string, in_err error) {
        
        state := STX
        l_tok := ""

        for _, ch := range i_line {
            //debug 
            //d fmt.Printf("%c", ch)
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
//x         l_running_len = 0
//x         l_rowTokens := NewRowTokens()
//x         l_rowTokens := nil
        var l_rowTokens *RowTokens = nil
        
//..    //x type Line struct {
//..    //x     tag         string     // typically filename
//..    //x     ln          int        // line number
//..    //x     line        string     // line contents
//..    //x }
//..    i_lines chan readers.Line

        for l_line := range i_lines {

            if l_running_len == 0 {
                l_rowTokens = NewRowTokens(l_line)
            }

            //x l_tokens, err := parse_line_v2(l_line)
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
                //x l_rowTokens = NewRowTokens()
            } else if  l_running_len > i_len {
                l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too many tokens - expected : `, i_len, " actual : ", l_running_len))
                output <- *l_rowTokens
                break
            }

            ix++
        }

        if l_running_len > 0 && l_running_len < i_len {
            l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too few tokens - expected : `, i_len, " actual : ", l_running_len))
            output <- *l_rowTokens
        }

    }()

    return output
}
