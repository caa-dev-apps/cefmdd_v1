package main

import (
     "log"
     "testing"
)

//- Keywords
//- Key,                    low,    high,   type,               scope,  source,     table,  comments__20160617,,
//- MISSION                 ,1,     1,      Enumerated         ,m,      CAA        ,3,      ,       ,
//- PARENT_MISSION          ,1,     1,      Enumerated         ,,       CAA        ,3,,,
//- MISSION_TIME_SPAN       ,1,     1,      ISO_TIME_RANGE     ,m,      CAA        ,3,,,

//- Enum
//- Keyword,                Definition,     Description,                                            Additional
//- MISSION,                Cluster,        ,
//- MISSION,                DoubleStar,     ,
//- PARENT_MISSION,         Cluster,        Mission or project from which the data was collected:,
//- MISSION_AGENCY,         ESA,            European Space Agency,

//-                 Key,                                    low,    high,      type,              
//-                 MISSION,                                1,      1,         Enumerated,        
//-                 PARENT_MISSION,                         1,      1,         Enumerated,        
//-                 MISSION_TIME_SPAN,                      1,      1,         ISO_TIME_RANGE,    
//-                 MISSION_AGENCY,                         1,      1,         Enumerated,        
//-                 MISSION_DESCRIPTION,                    1,      1,         String,            
//-                 MISSION_KEY_PERSONNEL,                  1,      N,         String,            
//-                 MISSION_REFERENCES,                     1,      N,         String,            
//-                 MISSION_REGION,                         1,      N,         Enumerated,        
//-                 MISSION_CAVEATS,                        0,      N,         String,            
//-                 OBSERVATORY,                            1,      1,         Enumerated,        
//-                 PARENT_OBSERVATORY,                     0,      1,         Enumerated,        
//-                 OBSERVATORY_DESCRIPTION,                1,      1,         String,            
//-                 OBSERVATORY_TIME_SPAN,                  1,      1,         ISO_TIME_RANGE,    
//-                 OBSERVATORY_REGION,                     1,      N,         Enumerated,        
//-                 OBSERVATORY_CAVEATS,                    0,      N,         String,            
//-                 EXPERIMENT,                             1,      1,         Enumerated,        
//-                 PARENT_EXPERIMENT,                      0,      1,         Enumerated,        
//-                 EXPERIMENT_DESCRIPTION,                 1,      1,         String,            
//-                 EXPERIMENT_REFERENCES,                  1,      N,         String,            
//-                 INVESTIGATOR_COORDINATES,               1,      1,         String,            
//-                 EXPERIMENT_KEY_PERSONNEL,               1,      N,         String,            
//-                 EXPERIMENT_CAVEATS,                     0,      N,         String,            
//-                 INSTRUMENT_NAME,                        1,      1,         Enumerated,        
//-                 PARENT_INSTRUMENT,                      1,      1,         Enumerated,        
//-                 INSTRUMENT_DESCRIPTION,                 1,      1,         String,            
//-                 INSTRUMENT_TYPE,                        1,      N,         Enumerated,        
//-                 MEASUREMENT_TYPE,                       1,      N,         Enumerated,        
//-                 INSTRUMENT_CAVEATS,                     0,      N,         String,            
//-                 DATASET_ID,                             1,      1,         Enumerated,        
//-                 PARENT_DATASET_ID,                      1,      1,         String,            
//-                 DATASET_TITLE,                          1,      1,         String,            
//-                 DATA_TYPE,                              1,      1,         Formatted,         
//-                 DATASET_DESCRIPTION,                    1,      1,         String,            
//-                 CONTACT_COORDINATES,                    1,      1,         String,            
//-                 TIME_RESOLUTION,                        1,      1*,        Numerical,         
//-                 MIN_TIME_RESOLUTION,                    1,      1*,        Numerical,         
//-                 MAX_TIME_RESOLUTION,                    1,      1*,        Numerical,         
//-                 PROCESSING_LEVEL,                       1,      1,         Enumerated,        
//-                 ACKNOWLEDGEMENT,                        1,      1,         String,            
//-                 DATASET_CAVEATS,                        0,      1,         String,            
//-                 PARAMETER_ID,                           1,      1,         String,            
//-                 PARAMETER_TYPE,                         1,      1,         Enumerated,        
//-                 CATDESC,                                1,      1,         String,            
//-                 ENTITY,                                 1,      1*,        Enumerated,        
//-                 PROPERTY,                               1,      1*,        Enumerated,        
//-                 FLUCTUATIONS,                           0,      1,         Enumerated,        
//-                 ERROR_PLUS,                             0,      1,         Numerical,         
//-                 ERROR_MINUS,                            0,      1,         Numerical,         
//-                 COMPOUND,                               0,      1,         Enumerated,        
//-                 COMPOUND_DEF,                           0,      1,         String,            
//-                 UNITS,                                  1,      1,         String,            
//-                 SI_CONVERSION,                          1,      1,         Enumerated,        
//-                 TENSOR_ORDER,                           0,      1,         Enumerated,        
//-                 TENSOR_FRAME,                           0,      1,         Enumerated,        
//-                 TARGET_SYSTEM,                          0,      1,         Enumerated,        
//-                 COORDINATE_SYSTEM,                      ?,      1,         Enumerated,        
//-                 FRAME_ORIGIN,                           0,      1,         Enumerated,        
//-                 FRAME_VELOCITY,                         0,      1,         Enumerated,        
//-                 FRAME,                                  ?,      1,         String,            
//-                 REPRESENTATION,                         ?,      1,         Formatted,         
//-                 SIZES,                                  ?,      N,         Formatted,         
//-                 DEPEND_0,                               0,      1,         String,            
//-                 DEPEND_i,                               0,      1,         String,            
//-                 DATA_TYPE??,                            0,      N,         Formatted,         
//-                 LABEL_i,                                0,      1,         String,            
//-                 DELTA_PLUS,                             0,      1,         Numerical,         
//-                 DELTA_MINUS,                            0,      1,         Numerical,         
//-                 VALUE_TYPE,                             1,      1,         Enumerated,        
//-                 SIGNIFICANT_DIGITS,                     1,      1*,        Integer,           
//-                 FILLVAL,                                1,      1,         Numerical,         
//-                 QUALITY,                                1,      1*,        Integer,           
//-                 PARAMETER_CAVEATS,                      0,      1,         String,            
//-                 FIELDNAM,                               0,      1,         String,            
//-                 LABLAXIS,                               0,      1,         String,            
//-                 SCALEMIN,                               0,      1,         Numerical,         
//-                 SCALEMAX,                               0,      1,         Numerical,         
//-                 SCALETYP,                               0,      1,         Enumerated,        
//-                 DISPLAYTYPE,                            0,      1,         Enumerated,        
//-                 LOGICAL_FILE_ID,                        1,      1,         String,            
//-                 PARENT_DATASET,                         1,      1,         String,            
//-                 VERSION_NUMBER,                         1,      1,         Integer,           
//-                 DATASET_VERSION,                        1,      1,         Text,              
//-                 FILE_TYPE,                              1,      1,         Enumerated,        
//-                 FILE_TIME_SPAN,                         1,      1,         ISO_TIME_RANGE,    
//-                 GENERATION_DATE,                        1,      1,         ISO_TIME,          
//-                 INGESTION_DATE,                         1,      1,         ISO_TIME,          
//-                 FILE_SIZE,                              1,      1,         Numerical,         
//-                 METADATA_TYPE,                          0,      1,         Enumerated,        
//-                 METADATA_VERSION,                       0,      1,         Text,              
//-                 FILE_CAVEATS,                           0,      1,         String,            

//-                 MISSION_TIME_SPAN,                      1,      1,         ISO_TIME_RANGE, 
//-                 
//-                 MISSION_REGION,                         1,      N,         Enumerated, 
//-                 FILE_TYPE,                              1,      1,         Enumerated,       
//-                 
//-                 GENERATION_DATE,                        1,      1,         ISO_TIME,   
//-                 FILE_SIZE,                              1,      1,         Numerical,         
//-                 
//-                 LOGICAL_FILE_ID,                        1,      1,         String,            
//-                 METADATA_VERSION,                       0,      1,         Text,              
//-                 
//-                 SIGNIFICANT_DIGITS,                     1,      1*,        Integer,           

/////////////////////////////////////////////////////////////////////////////// -------------------------------------------------------------------------------
//

var (
    mdd_data = NewMddData()
)


func Test_mdd_all_should_pass(t *testing.T) {

    l_test_data := []KeyVal {
        { key: "MISSION_REGION",              val: []string{ "Magnetosheath"}},                                    //    1,      N,         Enumerated, 
        { key: "FILE_TYPE",                   val: []string{ "cdf"}},                                              //    1,      1,         Enumerated,       
        { key: "GENERATION_DATE",             val: []string{ "2011-10-09T00:00:00Z" }},                            //    1,      1,         ISO_TIME,   
        { key: "MISSION_TIME_SPAN",           val: []string{ "2011-10-09T00:00:00Z/2011-10-10T00:00:00Z" }},       //    1,      1,         ISO_TIME_RANGE, 
        { key: "FILE_SIZE",                   val: []string{ "12345" }},                                           //    1,      1,         Numerical,         
        { key: "LOGICAL_FILE_ID",             val: []string{ "ANY STRING" }},                                      //    1,      1,         String,            
        { key: "METADATA_VERSION",            val: []string{ "ANY TEXT" }},                                        //    0,      1,         Text,              
        { key: "SIGNIFICANT_DIGITS",          val: []string{ "12345678"}},                                         //    1,      1*,        Integer,           
                                                              
        { key: "PARAMETER_TYPE",              val: []string{ "Support_Data"}},                                     //    1,      1*,        Integer,           
    }

    for _, d := range l_test_data {
        err := mdd_data.test_input(&d)
        if err != nil {
            log.Print(err)        
        }
    }
}

func Test_mdd_all_should_fail(t *testing.T) {

    l_test_data := []KeyVal {
        { key: "MISSION_REGION",              val: []string{ "XMagnetosheath"}},                                    //    1,      N,         Enumerated, 
        { key: "FILE_TYPE",                   val: []string{ "Xcdf"}},                                              //    1,      1,         Enumerated,       
        { key: "GENERATION_DATE",             val: []string{ "X2011-10-09T00:00:00Z" }},                            //    1,      1,         ISO_TIME,   
        { key: "MISSION_TIME_SPAN",           val: []string{ "X2011-10-09T00:00:00Z/2011-10-10T00:00:00Z" }},       //    1,      1,         ISO_TIME_RANGE, 
        { key: "FILE_SIZE",                   val: []string{ "X12345" }},                                           //    1,      1,         Numerical,         
        { key: "LOGICAL_FILE_ID",             val: []string{ "" }},                                                 //    1,      1,         String,            
        { key: "METADATA_VERSION",            val: []string{ "" }},                                                 //    0,      1,         Text,              
        { key: "SIGNIFICANT_DIGITS",          val: []string{ "X12345678"}},                                         //    1,      1*,        Integer,           
    }

    for _, d := range l_test_data {
        err := mdd_data.test_input(&d)
        if err != nil {
            log.Print(err)        
        }
    }
}







func Test_mdd_enums(t *testing.T) {

    println("X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X ");

    l_test_data := []KeyVal {
        { key: "PARAMETER_TYPE",              val: []string{ "Support_Data"}},                                     //    1,      1*,        Integer,           
    }

    for _, d := range l_test_data {
        err := mdd_data.test_input(&d)
        if err != nil {
            log.Print(err)        
        }
    }
    
    println("X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X X - X ");
    
}

