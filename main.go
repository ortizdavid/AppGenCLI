package main

import (
	"fmt"
	"os"
	"github.com/ortizdavid/AppGenCore/generators"
	"github.com/ortizdavid/AppGenCore/helpers"
)


func getCommands() []string {
	commands := [] string {"-name", "-lang", "-type", "-db"}
	return commands
}

func main() {
	var cliArgs = os.Args;
	var lenArgs = len(cliArgs)
	var application  *generators.ApplicationGenerator
	
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

		if helpers.ContainsMultiple(getCommands(), appCommand, langCommand, typeCommand, dbCommand) == false {
			fmt.Println(helpers.INVALID_COMMAND)
		} else {
			application.Generate(appName, lang, appType, db)
		}
		
	}
}