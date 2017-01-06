package readers

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"path"    
    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

type Line struct {
    tag         string     // typically filename
    ln          int        // line number
    line        string     // line contents
}

func (l Line) String() string {
//x      return fmt.Sprintf("{ln: %d  tag: %s  line %s}", l.ln, l.tag, strings.TrimSpace(l.line))
     return fmt.Sprintf("ln: %d\n  tag: %s\n  line: %s\n", l.ln, l.tag, strings.TrimSpace(l.line))
}


func EachLine(i_path string) chan Line {
	output := make(chan Line, 16)
    l_tag := path.Base(i_path)
    
	go func() {
		defer close(output)

		fmt.Println(diag.BoldMagenta("File Open"), i_path)

		fi, err := os.Open(i_path)
		if err != nil {
			return
		}
		//x defer fmt.Println(diag.BoldMagenta("File Close"), i_path)
		defer fi.Close()

		var reader *bufio.Reader
		if strings.HasSuffix(i_path, "gz") == false {
			reader = bufio.NewReader(fi)
		} else {
			fz, err := gzip.NewReader(fi)
			if err != nil {
				return
			}
			defer fz.Close()

			reader = bufio.NewReader(fz)
		}

        ln  := 0
        
		for {
			line, err := reader.ReadString('\n')
			output <- Line{tag: l_tag, ln: ln, line: line }
			if err == io.EOF {
				break
			}
            
            ln++
		}
	}()

	return output
}

