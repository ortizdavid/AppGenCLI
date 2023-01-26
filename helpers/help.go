package helpers

import (
	"fmt"
)

func BasicInfo() {
	fmt.Println("BASIC INFO:")
	fmt.Println("\tappgen - tool for generate applications")
	fmt.Println("\tSource code\thttps://www.github.com/ortizdavid/appgenerator")
	fmt.Println("\tVersion\t\t1.0.0")
	fmt.Println()
}

func Description() {
	fmt.Println("DESCRIPTION:")
	fmt.Println("\tThis tool helps create the structure of an application, including database\n"+
				"\tFor now it generates for Python and PHP")
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
	fmt.Println("\tlist-langs\tList all supported languages and applications")
	fmt.Println("\thelp\t\tShows helps for appgen")
	fmt.Println("\t-name\t\tProject Name")
	fmt.Println("\t-lang\t\tProgramming Language")
	fmt.Println("\t-type\t\tType of application (mvc, api)")
	fmt.Println("\t-db\t\tdatabase ")
	fmt.Println()
}

func Usage() {
	fmt.Println("USAGE:")
	fmt.Println("\tappgen -name <app name> -lang <language> -type <type of project> -db <database or dbms>")
	fmt.Println()
}

func Examples() {
	fmt.Println("EXAMPLES:")
	fmt.Println("\tappgen -name PythonWebMVC -lang python -type mvc -db mysql\t\tCreates a MVC App with Python and MySQL")
	fmt.Println("\tappgen -name PythpnAPI -lang python -type api -db postgres\t\tCreates an API with Python and Postgres")
	fmt.Println("\tappgen -name PHPSimpleApp -lang php -type mvc -db mysql\t\t\tCreates a MVC App with PHP and MySQL")
	fmt.Println("\tappgen help\t\t\t\t\t\t\t\tShows help comands")
	fmt.Println("\tappgen list-langs\t\t\t\t\t\t\tLists all suportded languages")
	fmt.Println()
}

func PrintHelp()  {
	BasicInfo()
	Description()
	Usage()
	Commands()
	Examples()
	Author()
}

func ListLanguages() {
	fmt.Println("SUPORTED LANGUAGES:")
	fmt.Println("\tLang\tprojects")
	fmt.Println("\tpython\tmvc, api")
	fmt.Println("\tphp\tmvc")
	fmt.Println()
}
