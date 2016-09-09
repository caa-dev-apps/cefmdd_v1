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

## Misc        
go test -run name-of-test-pattern  