# ~/C/CAA-DEV-APPS-SAMPLE-DATA/CEF_CEH_EXAMPLES_2013_VALIDATOR/
# /home/spw/C/CAA-DEV-APPS-SAMPLE-DATA/CEF_CEH_EXAMPLES_2013_VALIDATOR/CEF

ROOT=~/C/CAA-DEV-APPS-SAMPLE-DATA
ROOT_2013=$ROOT/CEF_CEH_EXAMPLES_2013_VALIDATOR

INC_0=$ROOT_2013/header_includes
INC_1=$ROOT_2013/header_includes/whisper
TEST_FILE_0=$ROOT_2013/CEF/WHISPER/C1_CP_WHI_ELECTRON_DENSITY__20080704_V01.cef.gz

cefmdd_v1 -i $INC_0 -i $INC_1 -f $TEST_FILE_0 $1 $2 $3 $4