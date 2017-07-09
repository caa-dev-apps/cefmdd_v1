# ~/C/CAA-DEV-APPS-SAMPLE-DATA/T/CEF_CEH_EXAMPLES_2013_VALIDATOR/
# /home/spw/C/CAA-DEV-APPS-SAMPLE-DATA/T/CEF_CEH_EXAMPLES_2013_VALIDATOR/CEF

ROOT=~/C/CAA-DEV-APPS-SAMPLE-DATA/T
ROOT_2013=$ROOT/CEF_CEH_EXAMPLES_2013_VALIDATOR

INC_0=$ROOT_2013/header_includes
INC_1=$ROOT_2013/header_includes/whisper
TEST_FILE_1=$ROOT_2013/CEF/WHISPER/C2_CP_WHI_SOUNDING_TIMES__20010203_V02.cef.gz

cefmdd_v1 -i $INC_0 -i $INC_1 -f $TEST_FILE_1 $1 $2 $3 $4