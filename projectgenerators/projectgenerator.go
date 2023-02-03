package projectgenerators

import (
	"fmt"
	"github.com/ortizdavid/appgen/helpers"
)

type IGenerator interface {
	GenerateConfig(rootDir string, db string)
	GenerateMain(rootDir string, appType string)
	GenerateModels(rootDir string)
	GenerateViews(rootDir string)
	GenerateStaticFiles(rootDir string)
	GenerateApiControllers(rootDir string)
	GenerateMvcControllers(rootDir string)
	GenerateMySqlDB(rootDir string)
	GeneratePostgresDB(rootDir string)
	InstallDeps(rootDir string)
	CreateApp(appName string, appType string, db string)
}

type ProjectGenerator struct {}

func (project *ProjectGenerator) Generate(appName string, lang string, appType string, db string)  {
	fmt.Printf(helpers.CREATING_PROJECT, appName)

	switch lang {
	case "python":
		var python *PythonProject
		python.CreateApp(appName, appType, db)
	}
}

func (project *ProjectGenerator) GetCommands() []string {
	commands := [] string {"-name", "-lang", "-type", "-db"}
	return commands
}

func (project *ProjectGenerator) GetTypes() []string {
	types := [] string {"api", "mvc"}
	return types
}

func (project *ProjectGenerator) GetDatabases() []string {
	databases := [] string {"postgres", "mysql"}
	return databases
}

func (project *ProjectGenerator) GetLanguages() []string {
	languages := [] string {"python"}
	return languages
}
