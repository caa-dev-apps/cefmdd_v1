ROOT=~/C/CAA-DEV-APPS-SAMPLE-DATA/T/TESTS_2017_FGM

INC_0=$ROOT/header_includes
INC_1=$INC_0/fgm

TEST_FILE=$ROOT/fgm/C1_CP_FGM_5VPS__20130602_192530_20130605_014110_V01.cef.gz


cefmdd_v1 -i $INC_0 -i $INC_1 -f $ROOT/fgm/C2_CQ_FGM_CAVF__20130602_192530_20130605_014110_V01.cef.gz  $1 $2 $3 $4
