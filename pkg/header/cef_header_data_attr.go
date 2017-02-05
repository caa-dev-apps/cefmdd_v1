package header

import (
	"github.com/caa-dev-apps/cefmdd_v1/pkg/diag"
	"github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
)

func (hds *CefHeaderData) kv_attr(kv *readers.KeyVal) (err error) {

    //- FILE_NAME
    //- FILE_FORMAT_VERSION
    //- END_OF_RECORD_MARKER
    //- DATA_UNTIL

    //- todo need to  add kv_var to current Var{}
    //- println("TODO: kv_attr");
    diag.Todo("kv_attr")
	return
}
