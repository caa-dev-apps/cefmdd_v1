@echo off

set XX=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_

@echo on

cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef.gz $1

pause

