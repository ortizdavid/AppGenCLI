package helpers

import (
	"fmt"
)

func Description() {
	fmt.Println("DESCRIPTION:")
	fmt.Println("\tappgen - tool for generate applications")
	fmt.Println("\tSource code\thttps://www.github.com/ortizdavid/appgenerator")
	fmt.Println("\tThis tool helps create the structure of an application, including database\n"+
				"\tFor now it generates for Python and PHP")
	fmt.Println("\tVersion\t\t1.0.0")
	fmt.Println()
}

func Author() {
	fmt.Println("AUTHOR:")
	fmt.Println("\tName\t\tOrtiz de Arcanjo António David")
	fmt.Println("\tPhone\t\t+244 936 166 699")
	fmt.Println("\tEmail\t\tortizaad1994@gmail.com")
	fmt.Println("\tGithub\t\thttps://www.github.com/ortizdavid")
	fmt.Println("\tLinkedIn\thttps://www.linkedin.com/in/ortiz-david")
	fmt.Println()
}

func Commands() {
	fmt.Println("COMANDS:")
	fmt.Println("\tappgen\t\tFirst command")
	fmt.Println("\thelp\t\tShows helps for appgen")
	fmt.Println("\texamples\tShows Examples")
	fmt.Println("\tusage\t\tShows Usage of appgen")
	fmt.Println("\tauthor\t\tShows all author information")
	fmt.Println("\tlist-langs\tList all supported languages and applications")
	fmt.Println("\t-name\t\tProject Name")
	fmt.Println("\t-lang\t\tProgramming Language")
	fmt.Println("\t-type\t\tType of application (mvc, api)")
	fmt.Println("\t-db\t\tdatabase ")
	fmt.Println()
}

func Usage() {
	fmt.Println("USAGE:")
	fmt.Println("\tFor create an Application:")
	fmt.Println("\tappgen -name <app name> -lang <language> -type <type of project> -db <database or dbms>")
	fmt.Println("\tFor Help and Examples:")
	fmt.Println("\t\tappgen <command>")
	fmt.Println()
}

func Examples() {
	fmt.Println("EXAMPLES:")
	fmt.Println("\tappgen -name PythonMVC -lang python -type mvc -db mysql\t\t\tCreates a MVC App with Python and MySQL")
	fmt.Println("\tappgen -name PythonAPI -lang python -type api -db postgres\t\tCreates an API with Python and Postgres")
	fmt.Println("\tappgen help\t\t\t\t\t\t\t\tShows help comands")
	fmt.Println("\tappgen list-langs\t\t\t\t\t\t\tLists all suportded languages")
	fmt.Println()
}

func PrintHelp()  {
	Description()
	Usage()
	Commands()
	Examples()
}

func ListLanguages() {
	fmt.Println("\nSUPORTED LANGUAGES:")
	fmt.Println("\tLang\tprojects")
	fmt.Println("\tpython\tmvc, api")
	fmt.Println()
}