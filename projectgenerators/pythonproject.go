package projectgenerators

import (
	"github.com/ortizdavid/appgen/filemanager"
	dbsamples "github.com/ortizdavid/appgen/samples/db"
	pythonsamples "github.com/ortizdavid/appgen/samples/python"
	"github.com/ortizdavid/appgen/samples/scripts"
	//dbsamples "github.com/ortizdavid/appgen/samples/db"
)

type PythonProject struct {}


var fm *filemanager.FileManager
var config *pythonsamples.ConfigPy
var appImport *pythonsamples.AppImport


func (python *PythonProject)  GenerateConfig(rootDir string, db string) {
	file := "config.py"
	fm.CreateFile(rootDir, file)
	fm.WriteFile(rootDir, file, appImport.ImportForConfig())
	fm.WriteFile(rootDir, file, config.CreateConfig(db))
}

func (python *PythonProject)  GenerateMain(rootDir string, appType string) {
	file := "main.py"
	fm.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		fm.WriteFile(rootDir, file, appImport.ImportForMvcApp())
	case  "api":
		fm.WriteFile(rootDir, file, appImport.ImportForRestApi())
	}
	fm.WriteFile(rootDir, file, appImport.AppMainCode())
}

func (python *PythonProject)  GenerateModels(rootDir string) {
	fm.CreateFolderAll(rootDir+"/models")
}

func (python *PythonProject)  GenerateViews(rootDir string) {
	fm.CreateFolderAll(rootDir+"/templates")
}

func (python *PythonProject)  GenerateStaticFiles(rootDir string) {
	fm.CreateFolderAll(rootDir+"/static/css")
	fm.CreateFolderAll(rootDir+"/static/js")
	fm.CreateFolderAll(rootDir+"/static/imgs")
}

func (python *PythonProject)  GenerateMvcControllers(rootDir string) {
	fm.CreateFolderAll(rootDir+"/controllers")
}

func (python *PythonProject)  GenerateApiControllers(rootDir string) {
	fm.CreateFolderAll(rootDir+"/api_controllers")
}

func (python *PythonProject)  GenerateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	fm.CreateFolderAll(dbDir)
	fm.CreateFile(dbDir, file)
	fm.WriteFile(dbDir, file, mysql.CreateDatabase())
	fm.WriteFile(dbDir, file, mysql.CreateTables())
	fm.WriteFile(dbDir, file, mysql.CreateSQLViews())
	fm.WriteFile(dbDir, file, mysql.InsertRoles())
	fm.WriteFile(dbDir, file, mysql.InsertUsers())
	fm.WriteFile(dbDir, file, mysql.InsertTasks())
}

func (python *PythonProject)  GeneratePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	fm.CreateFolderAll(dbDir)
	fm.CreateFile(dbDir, file)
	fm.WriteFile(dbDir, file, postgres.CreateDatabase())
	fm.WriteFile(dbDir, file, postgres.CreateTables())
	fm.WriteFile(dbDir, file, postgres.CreateSQLViews())
	fm.WriteFile(dbDir, file, postgres.InsertRoles())
	fm.WriteFile(dbDir, file, postgres.InsertUsers())
	fm.WriteFile(dbDir, file, postgres.InsertTasks())
}

func (python *PythonProject)  GenerateRequirements(rootDir string, db string) {
	var deps *scripts.PythonDeps
	file := "requirements.txt"
	fm.CreateFile(rootDir, file)
	fm.WriteFile(rootDir, file, deps.Requirements(db))
}


func (python *PythonProject) CreateApp(appName string, appType string, db string) {
	
	python.GenerateModels(appName)
	
	switch appType {
	case "mvc":
		python.GenerateViews(appName)
		python.GenerateStaticFiles(appName)
		python.GenerateMvcControllers(appName)

	case "api":
		python.GenerateApiControllers(appName)
	}

	switch db {
	case "mysql":
		python.GenerateMySqlDB(appName)
	case "postgres":
		python.GeneratePostgresDB(appName)
	}
	python.GenerateConfig(appName, db)
	python.GenerateRequirements(appName, db)
	python.GenerateMain(appName, appType)
}

