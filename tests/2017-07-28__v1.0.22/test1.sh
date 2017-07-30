#!/bin/sh

# /home/caa_ops/DEVELOPMENT/SPW/LATEST/cefmdd_v1  -i ./HEADERS/ -f ./CEF/C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V10.cef.gz
cefmdd_v1  -i ./HEADERS/ -f ./CEF/C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V10.cef.gz

if [ "$?" -eq "0" ]
then
	echo ""
	echo "ERROR: cefmdd returns 0 - this file is invalid, it should return -1"
	echo ""
else
	echo ""
	echo "SUCCESS: cefmdd returns non-0 - this fle is invalid, this is correct"
	echo ""
fi



# /home/caa_ops/DEVELOPMENT/SPW/LATEST/cefmdd_v1  -i ./HEADERS/ -f ./CEF/C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V01.cef.gz
cefmdd_v1  -i ./HEADERS/ -f ./CEF/C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V01.cef.gz

if [ "$?" -eq "0" ]
then
	echo ""
	echo "SUCCESS: cefmdd returns 0 - this file is valid, this is correct"
	echo ""
else
	echo ""
	echo "ERROR: cefmdd retunrs non-0 - this file is valid, it should return 0"
	echo ""
fi


