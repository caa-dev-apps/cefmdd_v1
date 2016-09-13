## cefmdd_v1
A cef meta data validator


The software is based on the rules provided by the latest Meta Data Dictionary (MDD) documentation.
An Excel workbook containing the canonical data for the Keywords and Enumertor types has been
derived from the MDD documentation - and this is currenlty read by the software at startup 
and then used to validate the cef/ceh files under test.

[Canonical MDD](https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/pubhtml "Google Hosted Excel Workbook")


## Status
Under development


## Installation
go install github.com\caa-dev-apps\cefmdd_v1


## Config
During development, 2 of the worksheets (Keywords, Enums) from the cannonical Google hosted Excel workbook (linked to above)
are copied to the /home/user/.cefmdd_v1 folder as csv files and used by the cefmdd_v1 App on startup for its rules.


## Example test run
set INC_0=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/HEADERS

set INC_1=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/HEADERS/EDI

set TEST_FILE_4=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/EDI/C3_CP_EDI_QZC__20111021_V01.cef.gz


cefmdd_v1 -i %INC_0% -i %INC_1% -f %TEST_FILE_4% 


## TODO
    parser.go 
    func (h *CefHeaderData) add_kv(kv *KeyVal) (err error) 
        -> mdd

        
    Name of file on disk equals FILE_NAME	Within Double Quotes
    Dataset name matches LOGICAL_FILE_ID (DATASET must be start of logical file id)	
    Filename matches LOGICAL_FILE_ID + FILE_TYPE	
    VERSION_NUMBER must be integer	
    Version in filename /logical_file id must match VERSION_NUMBER	
    FILE_TYPE must match that given in the filename	
    END_OF_RECORD_MARKER is not manditory, defaults to New Line char	in double quotes
    FILE_VERSION must be integer	
    FILE_TIME_SPAN must be ISO time	
    FIME_TIME_SPAN start time must be before stop time	
    FILE_TIME_SPAN:  File Start Time must be <= First data record; File End Time must be >= TRUNCATE_TO_SECONDS(Last data record) *	* TRUNCATE to seconds UNLESS the time is 00:00:00
    DATA_UNTIL - End of data has to be in double quotes 	
    OBSERVATORY should match the LOGICAL_FILE_ID	
    DATA_TYPE should match LOGICAL_FILE_ID	
        
    QUALITY : if there is a link to another parameter - validate that this parameter is in this dataset (if indicated) OR if param is in another dataset, validate it is the correct format Cn_Cx_EXPER_NAME	
        
    Make sure ISO time seconds are < 60 (CIS problem)	
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
## Misc        
go test -run name-of-test-pattern  

GOOS=linux GOARCH=386 go build

set GOOS=linux&&set GOARCH=amd64&& go build -v .
set GOOS=linux&&set GOARCH=amd64&& go build -v a1.go
set GOOS=linux&&set GOARCH=386&& go build -v .



    ~/.cefmdd_v1
        Enums.csv
        Keywords.csv
        
    ~/cefmdd_v1
        /CEF_CEH_SAMPLES
        cefmdd_v1
        t4.sh

            INC_0=~/cefmdd_v1/CEF_CEH_SAMPLES/HEADERS
            INC_1=~/cefmdd_v1/CEF_CEH_SAMPLES/HEADERS/EDI
            TEST_FILE_4=~/cefmdd_v1/CEF_CEH_SAMPLES/CEF/EDI/C3_CP_EDI_QZC__20111021_V01.cef.gz

            ./cefmdd_v1 -i $INC_0 -i $INC_1 -f $TEST_FILE_4

            
# csv sheets            
    https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/export?format=csv&gid=791408233
    https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/export?format=csv&gid=1933632029


Bash/sh    
    wget –no-check-certificate -q -O – ‘https://docs.google.com/spreadsheet/ccc?key=0At2sqNEgxTf3dEt5SXBTemZZM1gzQy1vLVFNRnludHc&single=true&gid=0&output=txt’ | cut -f1,3

R    
    covarianceTableURL=”https://docs.google.com/spreadsheet/ccc?key=0At2sqNEgxTf3dEt5SXBTemZZM1gzQy1vLVFNRnludHc&single=true&gid=0&output=txt”
    require(RCurl)
    covarianceTable=read.table(textConnection(getURL(covarianceTableURL)))
            
            
            
            
            
            
            
            