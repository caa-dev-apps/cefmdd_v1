@echo off

set INC_0=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/header_includes
set INC_1=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/header_includes/whisper
set TEST_FILE_4=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/WHISPER/C1_CP_WHI_ELECTRON_DENSITY__20080704_V01.cef.gz

@echo on

cefmdd_v1 -i %INC_0% -i %INC_1% -f %TEST_FILE_5%
pause

