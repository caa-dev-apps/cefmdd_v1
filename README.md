# cefmdd_v1

A cef meta data validator


The software is based on the rules provided by the latest Meta Data Dictionary (MDD) documentation.
An Excel workbook containing the canonical data for the Keywords and Enumertor types has been
derived from the MDD documentation - and this is currenlty read by the software at startup 
and then used to validate the cef/ceh files under test.

[Canonical MDD](https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/pubhtml "Google Hosted Excel Workbook")


# Status

* Parses cef/ceh files for header data validity.
    * creates in memory representation (attrs, meta, vars) for fast referencing.
* Performs header data tests detailed in the above spreadsheet (more to be added over time)
* Reads Data Records 
    * maps variables to data-types vector taking into account SIZES of each variable (Except DATA variables)
    * handles single/multiple lines - with/without END_OF_RECORD_MARKER
    * checks for increasing date-time stamps 
    * checks for correct number of cells pre record
    * checks each cell for correct data-type


# Installation

```bash
go get github.com\caa-dev-apps\cefmdd_v1\cmd\cefmdd_v1
```

# Config

During development, 2 of the worksheets (Keywords, Enums) from the cannonical Google hosted Excel workbook (linked to above)
are copied to the /home/user/_cefmdd_v1 folder as csv files and used by the cefmdd_v1 App on startup for its rules.


# Example test run

```bash
set INC_0=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/HEADERS
set INC_1=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/HEADERS/EDI
set TEST_FILE_4=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/EDI/C3_CP_EDI_QZC__20111021_V01.cef.gz

cefmdd_v1 -i %INC_0% -i %INC_1% -f %TEST_FILE_4% 
```


        
# Misc Notes

```
    go test -run name-of-test-pattern  
    GOOS=linux GOARCH=386 go build
    set GOOS=linux&&set GOARCH=amd64&& go build -v .
    set GOOS=linux&&set GOARCH=amd64&& go build -v a1.go
    set GOOS=linux&&set GOARCH=386&& go build -v .

```

