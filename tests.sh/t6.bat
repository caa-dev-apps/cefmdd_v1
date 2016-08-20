@echo off

set TEST_HOME=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_

set INC_0=%TEST_HOME%/header_includes
set INC_1=%TEST_HOME%/header_includes/whisper

@echo on

cefmdd_v1.exe -i %INC_0% -i %INC_1% -f %TEST_HOME%/CEF/WHISPER/C1_CP_WHI_ELECTRON_DENSITY__20080704_V01.cef.gz

pause

