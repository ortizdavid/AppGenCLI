package helpers

func Contains(s []string, str string)  bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsMultiple(s [] string, param1, param2, param3, param4 string)  bool {
	return Contains(s, param1) == true && Contains(s, param2) == true && Contains(s, param3) == true && Contains(s, param4) == true 
}