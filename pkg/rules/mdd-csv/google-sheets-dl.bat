@echo off
rem curl  https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/export?format=csv&gid=791408233 -o  "mdd__20160617.xlsx - Tables.csv"
rem curl  https://docs.google.com/spreadsheets/d/1KSEQS-1ncG7tNt7PJRVbT_gGuFfyJfmWNHoKANsw5kI/export?format=csv&gid=1933632029 -o "mdd__20160617.xlsx - Enums.csv"

set DEST=%HOMEDRIVE%\\%HOMEPATH%\\.cefmdd_v1
md %DEST%

@echo on

copy "mdd__20160617.xlsx - Keywords.csv"  %DEST%\\"Keywords.csv"
copy "mdd__20160617.xlsx - Enums.csv"     %DEST%\\"Enums.csv"

copy "mdd__20160617.xlsx - Keywords.csv"  %DEST%\\"Keywords.csv.txt"
copy "mdd__20160617.xlsx - Enums.csv"     %DEST%\\"Enums.csv.txt"

      