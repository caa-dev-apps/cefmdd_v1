UPDATE: cefmdd shoudld only return 0 if the CEF files that was tested is correct. Any error encountered in the file should results in a non-0 return value (eg. -1)

the test1.sh script validates 2 files and checks if the cefmdd return value is 0:
- the first file is invalid (so cefmdd should return a non-0 value) -> C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V10.cef.gz
- the second file is valid (so cefmdd should return 0) -> C1_CP_PEA_PITCH_SPIN_DEFlux__20030101_V01.cef.gz 

