package readers

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

type KVState int

const (
	B4_KEY KVState = iota
	KEY
	B4_EQ
	B4_VAL
	VAL_Q
	VAL
	B4_NEXT
	B4_NEXT_COMMENT
	B4_EOL
	EOL_COMMENT
	CONTINUE_NEXT_LINE
	COMMENT
	EOL
)

func state_str(s KVState) (str string) {

	switch s {
	case B4_KEY:                return "B4_KEY              "
	case KEY:                   return "KEY                 "
	case B4_EQ:                 return "B4_EQ               "
	case B4_VAL:                return "B4_VAL              "
	case VAL_Q:                 return "VAL_Q               "
	case VAL:                   return "VAL                 "
	case B4_NEXT:               return "B4_NEXT             "
	case B4_NEXT_COMMENT:       return "B4_NEXT_COMMENT     "
	case B4_EOL:                return "B4_EOL              "
	case EOL_COMMENT:           return "EOL_COMMENT         "
	case CONTINUE_NEXT_LINE:    return "CONTINUE_NEXT_LINE  "
	case COMMENT:               return "COMMENT             "
	case EOL:                   return "EOL                 "
	default:                    return "????????            "
	}

	return
}

func state_diag(ch rune, s1, s2 KVState) {

	diag.Println(string(ch), state_str(s1), state_str(s2))
}


type KeyVal struct {
	Key string
	Val []string
    
    Line Line       // 
}

func (kv KeyVal) String() string {
     return fmt.Sprintf("key: %s\n  val: %s\n  line: %s\n", kv.Key, kv.Val, kv.Line)
}

func (kv KeyVal) NewSwitchKey(k string) (*KeyVal) {
    return &KeyVal{Key: k, Val: kv.Val, Line: kv.Line}
}

func NewMetaKeyVal(k string, v []string) (*KeyVal) {
    return &KeyVal{Key: k, Val: v}
}


func EachKeyVal(i_lines chan Line) chan KeyVal{
	output := make(chan KeyVal, 16)

	state := B4_KEY
	key := ""
	val := []string{}

	l_key := ""
	l_val := ""

	initVars := func() {
		state = B4_KEY

		key = ""
		val = []string{}

		l_key = ""
		l_val = ""
	}

	parse_line := func(i_line string) (in_err error) {
	skip_comment:
		for _, ch := range i_line {
			//debug state_0 := state

			switch state {
			case B4_KEY:
				switch {
				case ch == '!':
					break skip_comment
				case unicode.IsSpace(ch):
					// No Change
				case unicode.IsLetter(ch):
					l_key += string(ch)
					state = KEY
				default:
					in_err = errors.New(`INVALID START OF LINE char -> ` + string(ch))
					return
				}
			case KEY:
				switch {
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-':
					l_key += string(ch)
				case unicode.IsSpace(ch):
					state = B4_EQ
				case ch == '=':
					key = l_key
					state = B4_VAL
				default:
					in_err = errors.New("INVALID KEY char -> " + string(ch))
					return
				}
			case B4_EQ:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '=':
					key = l_key
					state = B4_VAL
				default:
					in_err = errors.New("INVALID B4_EQ char -> " + string(ch))
					return
				}
			case B4_VAL:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '"':
					l_val = string(ch)
					state = VAL_Q
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+' || ch == '$':
					l_val = string(ch)
					state = VAL
				default:
					in_err = errors.New("INVALID B4_VAL char -> " + string(ch))
					return
				}
			case VAL_Q:
				switch ch {
				case '"':
					l_val += string(ch)
					val = append(val, l_val)
					state = B4_EOL
				default:
					l_val += string(ch)
				}
			case VAL:
				switch {
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+' || ch == ':' || ch == '/':
					l_val += string(ch)
				case unicode.IsSpace(ch):
					val = append(val, l_val)
					state = B4_EOL
				case ch == ',':
					val = append(val, l_val)
					state = B4_NEXT
				case ch == '!':
					val = append(val, l_val)
					state = EOL_COMMENT
				default:
					in_err = errors.New("INVALID VAL char -> " + string(ch))
					return
				}
			case B4_NEXT:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '\\':
					// No Change .. multiline
				case ch == '"':
					l_val = string(ch)
					state = VAL_Q
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+':
					l_val = string(ch)
					state = VAL
				case ch == '!':
					state = B4_NEXT_COMMENT
				default:
					in_err = errors.New("INVALID B4_NEXT char -> " + string(ch))
					return
				}

			case B4_NEXT_COMMENT:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '\\':
					// No Change .. multiline
					state = B4_NEXT
				}

			case B4_EOL:
				switch {
				case ch == '\\':
					state = CONTINUE_NEXT_LINE
				case unicode.IsSpace(ch):
					// No Change
				case ch == ',':
					state = B4_NEXT
				case ch == '!':
					break skip_comment
				default:
					in_err = errors.New("INVALID B4_EOL char -> " + string(ch))
					return
				}


			case CONTINUE_NEXT_LINE:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == ',':
					state = B4_NEXT
				case ch == '!':
					break skip_comment
				default:
					in_err = errors.New("INVALID CONTINUE_NEXT_LINE char -> " + string(ch))
					return
				}
			}





			//debug state_diag(ch , state_0, state)
		}

		if state == VAL {
			val = append(val, l_val)
			state = EOL
		}

		return
	}

	go func() {
		defer close(output)

		initVars()
        
		for l_line := range i_lines {

			err := parse_line(l_line.line)
			if err != nil {
                diag.Error("Tag->",    l_line.tag)
                diag.Error("Line #->", l_line.ln)
                diag.Error("Line->",   l_line.line)
				diag.Error("Error: ",  err)
				break
			}

			if (state != B4_NEXT) && (state != CONTINUE_NEXT_LINE) {
				// 2017-02-11 keys are case insensitive + meta names (handle on START_META)
				key = strings.ToUpper(key)

				if len(key) > 0 {
					output <- KeyVal{key, val, l_line}
				}

                if strings.EqualFold("DATA_UNTIL", key) == true {
                    break
                }
                
				initVars()
			}
		}
	}()

	return output
}
