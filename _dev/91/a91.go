package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	//x     "testing"
)

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
