
ON HOLD
=======


    func (h *CefHeaderData) check_mdd(kv *KeyVal) (err error) {

    switch {
        case kv.key == "ENTRY": return          // check meta or valiable etx
        case kv.key == "FILLVAL": return        // multiple types - depends on var type
        case kv.key == "SIZES": return          // type is FORMAT can be 1 or 1,2 for e.g.
        default:
    }



KEYWORDS.csv
============

patterns:
            DEPEND_i                    // removed DEPEND_0
            LABEL_i
            REPRESENTATION_i



from:       UNITS                       1, 1         
to:         UNITS                       1, N

from:       SI_CONVERSION               1, 1         
to:         SI_CONVERSION               1, N



from:       REPRESENTATION_i            formatted -> 0, 1
            LABEL_i                     string    -> 0, 1      
            DEPEND_i                    string    -> 0, 1
            
to:         REPRESENTATION_i            string    -> 0, N
            LABEL_i                     string    -> 0, N     
            DEPEND_i                    string    -> 0, N
            
         

from:       DATASET_ID                  ,1,1,       Enumerated
to:         DATASET_ID                  ,1,1,       string

     
from        PARENT_DATASET              ,1,1,       Enumerated      
to:         PARENT_DATASET              ,1,1,       String         
     

     
from:       MISSION_DESCRIPTION         ,1,1,       String
            OBSERVATORY_DESCRIPTION     ,1,1,       String
            EXPERIMENT_DESCRIPTION      ,1,1,       String
            INSTRUMENT_DESCRIPTION      ,1,1,       String
            DATASET_DESCRIPTION         ,1,1,       String
            INVESTIGATOR_COORDINATES    ,1,1,       String      
            DATASET_CAVEATS             ,0,1,       String
            
to:         MISSION_DESCRIPTION         ,1,N,       String
            OBSERVATORY_DESCRIPTION     ,1,N,       String
            EXPERIMENT_DESCRIPTION      ,1,N,       String
            INSTRUMENT_DESCRIPTION      ,1,N,       String
            DATASET_DESCRIPTION         ,1,N,       String
            INVESTIGATOR_COORDINATES    ,1,N,       String      
            DATASET_CAVEATS             ,0,N,       String
     
     
     
ENUMS.csv
=========

added:      SI_CONVERSION	            1.0>s	    FIXME																							
added:      SI_CONVERSION	            1>unitless	FIXME																							
added:      SI_CONVERSION	            1>degree  	FIXME																							


removed:    DATASET_ID


