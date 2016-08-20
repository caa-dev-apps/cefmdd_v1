@echo off

set XX=C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_

rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_PITCH_3DXL_DPFlux__20010101_V01.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_3DXH_DEFlux__20081231_V05.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_3DRH_DPFlux__20130817_V01.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/staff -f %XX%/CEF/STAFF/C2_CP_STA_AGC__20140825_V07.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/staff -f %XX%/CEF/STAFF/C1_CP_STA_SM__20140911_V07.cef.gz
rem cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/staff -f %XX%/CEF/STAFF/C1_CP_STA_AGC__20140911_V07.cef.gz

@echo on

cefmdd_v1 -i %XX%/header_includes -i %XX%/header_includes/peace -f %XX%/CEF/PEACE/C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef.gz

pause

