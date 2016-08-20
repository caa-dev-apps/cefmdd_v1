package main

import (
	"errors"
	"strings"
)


func (hds *CefHeaderData) kv_meta_entry(kv *KeyVal) (err error) {
    //todo need to  add entry to current Meta{}
    println("TODO: kv_meta_entry");
    return 
}

func (hds *CefHeaderData) kv_meta_value_type(kv *KeyVal) (err error) {
    //todo need to  add value_type to current Meta{}
    println("TODO: kv_meta_value_type");
	return
}

func (hds *CefHeaderData) kv_meta(kv *KeyVal) (err error) {

	switch {
	case strings.EqualFold("ENTRY", kv.key) == true:
		err = hds.kv_meta_entry(kv)
	case strings.EqualFold("VALUE_TYPE", kv.key) == true:
		err = hds.kv_meta_value_type(kv)
	default:
		return errors.New("META KEY: unknown")
	}

	return
}



// todo NEW for cef mdd_v1
func (hds *CefHeaderData) stx_meta(name *string) (err error) {
	//todo v1 hds.m_meta = Meta{}
	//todo v1 hds.m_meta.name = *name

	return
}

// todo NEW for cef mdd_v1
func (hds *CefHeaderData) etx_meta() (err error) {
	//todo v1 hds.m_data.metamap = append(hds.m_data.metamap, hds.m_meta)
    
	return
}











