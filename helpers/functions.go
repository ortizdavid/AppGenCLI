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
	return Contains(s, param1) && 
		Contains(s, param2) && 
		Contains(s, param3) && 
		Contains(s, param4)
}

func StrDatabase(db string) string {
	var database = ""
	switch db {
	case "mysql":
		database = "MySQL"
	case "postgres": 
		database = "PostgreSQL"
	}
	return database
}