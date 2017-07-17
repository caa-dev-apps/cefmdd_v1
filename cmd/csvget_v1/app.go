package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"time"

    "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

//  ## Sheet "A Test":
//  	                                                          https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1573083976&single=true&output=csv
//  	wget --no-check-certificate --output-document=a-test.csv 'https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1573083976&single=true&output=csv'
//  
//  
//  
//  ## Sheet "Keywords":
//  		                                                        https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=791408233&single=true&output=csv
//  	wget --no-check-certificate --output-document=keywords.csv 'https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=791408233&single=true&output=csv'
//  
//  ## Sheet "Enums":
//  	                                                         https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1933632029&single=true&output=csv
//  	wget --no-check-certificate --output-document=enums.csv 'https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1933632029&single=true&output=csv'

//x func main() {
//x 	res, err := http.Get("http://www.google.com/robots.txt")
//x 	if err != nil {
//x 		log.Fatal(err)
//x 	}
//x 	robots, err := ioutil.ReadAll(res.Body)
//x 	res.Body.Close()
//x 	if err != nil {
//x 		log.Fatal(err)
//x 	}
//x 	fmt.Printf("%s", robots)
//x }

// todo download f0
// todo download f1
// todo download f2
// todo 
// todo getDateTimeStamp - 20170716_hhmmss
// todo os.MkdirAll(home/.cefmdd_v1/stamp)
// todo 
// todo rename ~/.cefmdd_v1/atest.csv, stamp/atest.csv
// todo rename ~/.cefmdd_v1/atest.csv, stamp/keywords.csv
// todo rename ~/.cefmdd_v1/atest.csv, stamp/enums.csv
// todo 
// todo save f0 ~/.cefmdd_v1/atest.csv
// todo save f1 ~/.cefmdd_v1/atest.csv
// todo save f2 ~/.cefmdd_v1/atest.csv
// todo 

func nowString() (s string) {

	t := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d %02d.%02d.%02d",	
						t.Year(),
					    t.Month(),
					    t.Day(),
					    t.Hour(),
					    t.Minute(),
					    t.Second())
}


func httpGetFile(fn string) (contents string, err error) {

//	res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get(fn)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	contents = string(b)

	fmt.Printf("%s\n", contents)

	return
}


func main() {
    diag.Info(diag.Yellow("csvget v0.0.0, (16 July 2017)"))


	files := map[string]string {
		"Atest.csv": 	`https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1573083976&single=true&output=csv`,		// A Test
		"Keywords.csv": `https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=791408233&single=true&output=csv`,			// Keywords
		"Enums.csv": 	`https://docs.google.com/spreadsheets/d/1CFFyoD4pKRq0vyaEWm8pnX3hxrjEEfzG69f3j1PTuHw/pub?gid=1933632029&single=true&output=csv`,		// Enums
	}

	file_contents := make(map[string]string)

	for fn, url := range files {
		contents, err := httpGetFile(url)
		if err != nil {
			log.Fatal(err)		
		} 

		file_contents[fn] = contents
	}

	// s := fmt.Printf("Bak.%s", nowString())
	home := os.Getenv("HOME")
	tmp  := fmt.Sprintf("%s/_cefmdd_v1/%s Bak", 
						home,
					  	nowString())
	fmt.Println("", home)
	fmt.Println("", tmp)

	err := os.MkdirAll(tmp,  0755)
	if err != nil {
		log.Fatal(err)		
	} 




    diag.Info(diag.Yellow("CAA Rocks!\n"))
}