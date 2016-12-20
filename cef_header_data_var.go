package main

import ()

func (hds *CefHeaderData) kv_var(kv *KeyVal) (err error) {

    //todo need to  add kv_var to current Var{}
    println("TODO: kv_var");
	return
}

func (hds *CefHeaderData) stx_var(name *string) (err error) {
	//todo v1 hds.m_var = Var{}
	//todo v1 hds.m_var.name = *name    
    
	return
}

func (hds *CefHeaderData) etx_var() (err error) {
	//todo v1 hds.m_data.varlist = append(hds.m_data.varlist, hds.m_var)
    //todo v1 hds.m_data.varmap = append(hds.m_data.varmap, hds.m_var)
    
	return
}
