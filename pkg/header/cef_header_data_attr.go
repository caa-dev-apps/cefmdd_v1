package header

import (
	"github.com/caa-dev-apps/cefmdd_v1/pkg/readers"
)

func (hds *CefHeaderData) kv_attr(kv *readers.KeyVal) (err error) {

    //x FILE_NAME
    //x FILE_FORMAT_VERSION
    //x END_OF_RECORD_MARKER
    //x DATA_UNTIL

    //todo need to  add kv_var to current Var{}
    println("TODO: kv_attr");

	return
}
