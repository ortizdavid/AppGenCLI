package main

import (
	"fmt"
	"os"

	"github.com/ortizdavid/appgen/filemanager"
	"github.com/ortizdavid/appgen/helpers"

	pythonsamples "github.com/ortizdavid/appgen/samples/python"
)

var commands = []string {"-name", "-lang", "-type", "-db"}
var languages = []string {"php", "python"}
var types = []string {"api", "mvc"}
var databases = []string {"postgres", "mysql"}


func CreateApplication(appName string, lang string, appType string, db string)  {
	fmt.Printf(helpers.CREATING_PROJECT, appName)
	var fm *filemanager.FileManager
	fm.CreateFolder(appName)
}


func main() {

	var cliArgs = os.Args;
	var lenArgs = len(cliArgs)
	
	if lenArgs == 1 {
		helpers.PrintHelp()
	} else if lenArgs == 2 {
		var secondArg = cliArgs[1]
		if secondArg == "help" {
			helpers.PrintHelp()
		} else if secondArg == "list-langs" {
			helpers.ListLanguages()
		} else {
			return
		}
	} else if lenArgs == 9 {

		appCommand, appName := cliArgs[1], cliArgs[2]
		langCommand, lang := cliArgs[3], cliArgs[4]
		typeCommand, appType := cliArgs[5], cliArgs[6]
		dbCommand, db := cliArgs[7], cliArgs[8]

		if helpers.ContainsMultiple(commands, appCommand, langCommand, typeCommand, dbCommand) == false {
			fmt.Println(helpers.INVALID_COMMAND)

		} else if helpers.Contains(languages, lang) == false {
			fmt.Printf(helpers.UNSUPORTED_LANGUAGE, lang)

		} else if helpers.Contains(types, appType) == false {
			fmt.Printf(helpers.UNSUPORTED_TYPE, appType)

		} else if helpers.Contains(databases, db) == false {
			fmt.Printf(helpers.UNSUPORTED_DB, db)

		} else {
			CreateApplication(appName, lang, appType, db)
			var fm *filemanager.FileManager
			var config *pythonsamples.ConfigPy
			fm.CreateFile("flaskApi", "app.py")
			fmt.Println(config.GetConfig(db))
		}
	}
}

