INC_0=./HEADERS
INC_1=./HEADERS/efw
CEFMDD=/home/caa_ops/DEVELOPMENT/SPW/LATEST/cefmdd_v1

cefmdd_v1 -i $INC_0 -i $INC_1 -f ./CEF/C4_CP_EFW_L2_E3D_INERT_EX__20150101_000000_20150102_000000_V01.cef.gz
cefmdd_v1 -i $INC_0 -i $INC_1 -f C4_CP_EFW_L2_E3D_INERT_EX__20150101_000000_20150102_000000_V01.cef.gz

#$CEFMDD -i $INC_0 -i $INC_1 -f ./CEF/C4_CP_EFW_L2_E3D_INERT_EX__20150101_000000_20150102_000000_V01.cef.gz
#$CEFMDD -i $INC_0 -i $INC_1 -f C4_CP_EFW_L2_E3D_INERT_EX__20150101_000000_20150102_000000_V01.cef.gz
