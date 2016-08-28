# cefmdd_v1
A cef meta data validator


The software is based on the rules provided by the latest Meta Data Dictionary (MDD) documentation.
An Excel workbook containing the canonical data for the Keywords and Enumertor types has been
derived from the MDD documentation - and this is currenlty read by the software at startup 
and then used to validate the cef/ceh files under test.

[Canonical MDD](https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/pubhtml "Google Hosted Excel Workbook")


# Status
Under development


# Installation
go install github.com\caa-dev-apps\cefmdd_v1






# TODO
    parser.go 
    func (h *CefHeaderData) add_kv(kv *KeyVal) (err error) 
        -> mdd
        
        
        
  PARAMETER_TYPE error nil
  
  
go test -run name-of-test-pattern  