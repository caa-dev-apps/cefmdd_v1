package main

import (
	"fmt"
//x     "strings"
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
    COMMENT
    ERROR
    EOL
)

func state_str(s ROWState) (str string) {

    switch s {
    case STX:                   return "STX        "
    case TOK:                   return "TOK         "
    case TOK_2:                 return "TOK_2       "
    case QSTR:                  return "QSTR        "
    case QSTR_2:                return "QSTR_2      "
    case COMMENT:               return "COMMENT     "
    case ERROR:                 return "ERROR       "
    case EOL:                   return "EOL         "
    default:                    return "????????    "
    }

    return
}

func state_diag(ch rune, s1, s2 ROWState, tok string) {

    fmt.Println(string(ch), state_str(s1), state_str(s2), tok)
}


type RowTokens struct {
    Err error
    Tokens []string
    line string       // 
}

func (row RowTokens) String() string {
     return fmt.Sprintf("In: %s\n Tokens: %s\n", row.line, row.Tokens)
}

func (row *RowTokens) Append(t string)  {
     row.Tokens = append(row.Tokens, t)
}

func NewRowTokens() (*RowTokens) {
    return &RowTokens{
        Tokens: make([]string, 0),
    }
}



func rowTokeniser(i_lines chan string, i_len int) chan RowTokens {
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
                    // e.g. ... " 2002-11-10T04:17:14.758391Z,    -502,    -596,    -502,    -763, -10, -1000000000 $ "
                    fmt.Printf("________>> \tlen(r_tokens) %d \tl_running_len %d \t(i_len -1) %d \n", len(r_tokens), l_running_len, (i_len -1))
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
        }

        if state == TOK {
            r_tokens = append(r_tokens, l_tok)
        }

        return
    }


    go func() {
        defer close(output)

        ix := 0
        l_running_len = 0
        l_rowTokens := NewRowTokens()
        
        for l_line := range i_lines {
            //d fmt.Println("ix: -> ", ix, " ", l_line)

            l_tokens, err := parse_line_v2(l_line)
            l_rowTokens.Tokens = append(l_rowTokens.Tokens, l_tokens...)

            if err != nil {
                //d fmt.Println("Error: ",  err)
                //d fmt.Println("Line->",   l_line)

                l_rowTokens.Err = err
                output <- *l_rowTokens
                //x break
                return
            }

            l_running_len = len(l_rowTokens.Tokens)

            if l_running_len == i_len {
                output <- *l_rowTokens
                l_running_len = 0
                l_rowTokens = NewRowTokens()
            }

            if l_running_len > i_len {
                //d fmt.Println("Too many tokens", i_len, "actual", l_running_len)

                //x l_rowTokens.Err = errors.New(`error: line reader - Too many tokens`)
                l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too many tokens - expected : `, i_len, " actual : ", l_running_len))
                output <- *l_rowTokens
                break
            }

            ix++
        }

        //x if l_running_len < i_len {
        if l_running_len > 0 && l_running_len < i_len {
            //d fmt.Println("Too few tokens", i_len, "actual", l_running_len)

            l_rowTokens.Err = errors.New(fmt.Sprint(`error: line reader - Too few tokens - expected : `, i_len, " actual : ", l_running_len))
            output <- *l_rowTokens
        }


    }()

    return output
}

