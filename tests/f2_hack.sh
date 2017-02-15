ROOT=~/C/CAA-DEV-APPS-SAMPLE-DATA/T/TESTS_2017_FGM_HX1
INC_0=$ROOT/header_includes
INC_1=$INC_0/fgm

TEST_FILE=$ROOT/fgm/C1_PP_FGM_20010606_V01.cef
                    

cefmdd_v1 -i $INC_0 -i $INC_1 -f $TEST_FILE  $1 $2 $3 $4
