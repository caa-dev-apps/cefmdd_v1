package readers

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
  "github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
)

var test_cef_lines_txt_01 = `
!
    FILE_NAME           =   "C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef"
    INCLUDE             =   "C1_CH_PEA_PITCH_3DXLARH_DPFlux.ceh" ! Contains all the static meta-data
!
! Start of non-Static File Level - i.e. can't be from included file
!
    AAA                 =   1,2,3,4,5,      6   , 7 , 8 !HELLO THIS IS A COMMENT
    BBB                 =   0.0000E+00, 0.1099E+00, 0.2197E+00, 0.3296E+00,  !HELLO THIS IS A COMMENT \
                            0.4395E+00, 0.5493E+00, 0.6592E+00, 0.7691E+00  
    CCC                 =   "0 : SW-1>Solar Wind-Mode-1",\
                            "1 : SW-2>Solar Wind-Upstreams-Mode-2",\
                            "2 : SW-3>Solar Wind-Mode-3",\
                            "3 : SW-4>Solar Wind-Upstreams-Mode-4",\
                            "4 : SW-C1>Solar Wind-Data compression"
    DDD                 =   1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,\
                            33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64
    FFF                 =   "Re","Im"
    GGG                 =   22.500, 45.000, 45.000, 45.000, 45.000, 22.500
    REPRESENTATION_3    =   "x","y","z"
`

var test_cef_lines_txt_02 = `
    FILE_NAME="C3_CP_EDI_EGD__20111009_V01.cef"
    FILE_FORMAT_VERSION="CEF-2.0"
    ! FOLLOWING LINE WAS NOT IN THE ORIGIANL C3_CP_EDI_EGD__20111009_V01.cef.gz
    END_OF_RECORD_MARKER = ""
    !
    START_META = LOGICAL_FILE_ID
     ENTRY = "C3_CP_EDI_EGD__20111009_V01"
    END_META = LOGICAL_FILE_ID
    !
    START_META = VERSION_NUMBER
     ENTRY = "01"
    END_META = VERSION_NUMBER
    !
    START_META = FILE_TIME_SPAN
     VALUE_TYPE = ISO_TIME_RANGE
     ENTRY = 2011-10-09T00:00:00Z/2011-10-10T00:00:00Z
    END_META = FILE_TIME_SPAN
    !
    START_META = GENERATION_DATE
     VALUE_TYPE = ISO_TIME
     ENTRY = 2012-04-11T15:57:15Z
    END_META = GENERATION_DATE
    !
    START_META = FILE_CAVEATS
     ENTRY = "CAA_EDITOF_V1_04  2011-04-04T10:30:00Z"
    END_META = FILE_CAVEATS
    !
    ! include EGD (DATASET) HEADER File for Cluster-3
    ! with variable definitions, metadata_type and _version
    include="C3_CH_EDI_EGD_DATASET.ceh"
    !
    DATA_UNTIL=EOF
    !
`

func eachLineText(i_text string) chan Line {
    
	output := make(chan Line)

	go func() {
		defer close(output)
		scanner := bufio.NewScanner(strings.NewReader(i_text))
        
        l_tag := "TEST LineReader"
        l_ln := 0

        
		for scanner.Scan() {
			l := scanner.Text()
			if len(l) > 0 {
                output <- Line{tag: l_tag, ln: l_ln, line: l }
			}
            
            l_ln++
		}
	}()

	return output
}

func read_cef_chan(i_test_about string,
				   t *testing.T,
 				   i_lines chan Line) {
	cx := 0
	for _ = range eachKeyVal(i_lines) {
		cx++
	}

	if cx > 0 {
		t.Log(i_test_about, "cx = ", cx)
		diag.Println(i_test_about, "cx = ", cx)
	} else {
		t.Error(i_test_about, "cx = ", cx)
	}
}

func read_cef_text(i_test_name string,
                   t *testing.T,
                   i_text string) {
	read_cef_chan(i_test_name+" : Text",
                  t,
                  eachLineText(i_text))
}

func read_cef_file(i_test_name string,
                   t *testing.T,
                   i_filename string) {

	read_cef_chan(i_test_name+" : "+i_filename,
                  t,
                  EachLine(i_filename))
}

func TestRead_01_cef(t *testing.T) {
	read_cef_file("TestRead_01_cef",
				  t,
				  `C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/EDI/C3_CP_EDI_EGD__20111009_V01.cef`)
}

func TestRead_02_cef(t *testing.T) {
	read_cef_file("TestRead_02_cef",
				  t,
				  `C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/PEACE/C1_CP_PEA_PITCH_3DXH_DEFlux__20040421_V02.cef`)
}

func TestRead_03_cef_gz(t *testing.T) {
	read_cef_file("TestRead_03_cef_gz",
				  t,
				  `C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/EDI/C3_CP_EDI_EGD__20111009_V01.cef.gz`)
}

func TestRead_04_cef_gz(t *testing.T) {
	read_cef_file("TestRead_04_cef_gz",
				  t,
				  `C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_/CEF/PEACE/C1_CP_PEA_PITCH_3DXH_DEFlux__20040421_V02.cef.gz`)
}

func TestRead_10_cef_text(t *testing.T) {
	read_cef_text("TestRead_10_cef_text",
				  t,
				  test_cef_lines_txt_01)
}

func TestRead_11_cef_text(t *testing.T) {
	read_cef_text("TestRead_11_cef_text",
				  t,
				  test_cef_lines_txt_02)
}
