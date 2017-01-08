	fmt.Println(diag.BoldMagenta("cefmdd_v1 v0.0.5, (4 Jan 2016)"))
	fmt.Println(diag.BoldRed("Invalid Command Line Args"))	
	fmt.Println(diag.BoldRed(fmt.Sprintf("Error parsing cef file \n%#v", err.Error())))

	fmt.Println(diag.BoldYellow("cef cefpath"), a1s.m_cefpath)
	fmt.Println(diag.BoldRed(fmt.Sprintf("Error reading data record %#v", l_record.Line)))

	fmt.Println(diag.BoldRed(fmt.Sprintf("Error in data record, incorrect data type\n" +
										 "%s\n" +
										 "Error Token %s\n" +
										 "Tokens %#v\n" +
										 "Line %#v\n",
										 ct,
										 rd,
										 l_record.Tokens,
										 l_record.Line)))




    fmt.Println(v0_filename)
  	fmt.Println(k, v)     

  	fmt.Println(k, "Meta:ENTRY", entry)
  	fmt.Println("error:", err)
  	fmt.Println("--\n", h.m_vars)
	fmt.Println(diag.BoldRed(err.Error()), kv.Key, kv.Val)




	fmt.Println(string(ch), state_str(s1), state_str(s2))








= cef_reader_data
    - Add test for DATA_UNTIL
    - Add cmdline arg for max number of records
        - Default = 100


= tests
    copy all ceh files to dir for easy testing