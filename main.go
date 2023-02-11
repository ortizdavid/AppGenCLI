package main

import (
	"fmt"
	"os"
	"github.com/ortizdavid/appgen/projectgenerators"
	"github.com/ortizdavid/appgen/helpers"
)


func main() {
	var cliArgs = os.Args;
	var lenArgs = len(cliArgs)
	var project  *projectgenerators.ProjectGenerator
	
	if lenArgs == 1 {
		helpers.PrintHelp()

	} else if lenArgs == 2 {
		switch cliArgs[1] {
		case "help":
			helpers.PrintHelp()
		case "list-langs":
			helpers.ListLanguages()
		case "examples":
			helpers.Examples()
		default:
			helpers.Usage()
		}

	} else if lenArgs == 9 {

		appCommand, appName := cliArgs[1], cliArgs[2]
		langCommand, lang := cliArgs[3], cliArgs[4]
		typeCommand, appType := cliArgs[5], cliArgs[6]
		dbCommand, db := cliArgs[7], cliArgs[8]

		if helpers.ContainsMultiple(project.GetCommands(), appCommand, langCommand, typeCommand, dbCommand) == false {
			fmt.Println(helpers.INVALID_COMMAND)

		} else if helpers.Contains(project.GetLanguages(), lang) == false {
			fmt.Printf(helpers.UNSUPORTED_LANGUAGE, lang)

		} else if helpers.Contains(project.GetDatabases(), db) == false {
			fmt.Printf(helpers.UNSUPORTED_DB, db)

		} else {
			project.Generate(appName, lang, appType, db)
		}
		
	}
}