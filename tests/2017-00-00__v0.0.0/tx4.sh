INC_0=~/cefmdd_v1/CEF_CEH_SAMPLES/HEADERS
INC_1=~/cefmdd_v1/CEF_CEH_SAMPLES/HEADERS/EDI
TEST_FILE_4=~/cefmdd_v1/CEF_CEH_SAMPLES/CEF/EDI/C3_CP_EDI_QZC__20111021_V01.cef.gz

./cefmdd_v1 -i \$INC_0 -i \$INC_1 -f \$TEST_FILE_4  $1 $2 $3 $4

