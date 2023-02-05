package projectgenerators

import (
	"github.com/ortizdavid/appgen/filemanager"
	"github.com/ortizdavid/appgen/helpers"
	dbsamples "github.com/ortizdavid/appgen/samples/db"
	pythonsamples "github.com/ortizdavid/appgen/samples/python"
	"github.com/ortizdavid/appgen/samples/scripts"
)

type PythonProject struct {}

var fileManager *filemanager.FileManager
var readme *helpers.ReadMePy
var appImport *pythonsamples.AppImport
var config *pythonsamples.ConfigPy
var model *pythonsamples.Model
var mvcController *pythonsamples.MvcController
var apiController *pythonsamples.ApiController
var staticFile *pythonsamples.StaticFile

func (python *PythonProject) GenerateConfig(rootDir string, db string) {
	file := "config.py"
	fileManager.CreateFile(rootDir, file)
	fileManager.WriteFile(rootDir, file, appImport.ImportForConfig())
	fileManager.WriteFile(rootDir, file, config.CreateConfig(db))
}

func (python *PythonProject) GenerateReadme(rootDir string, db string, appType string) {
	file := "README.md"
	fileManager.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		fileManager.WriteFile(rootDir, file, readme.ReadmeMVC(db))
	case "api":
		fileManager.WriteFile(rootDir, file, readme.ReadmeAPI(db))
	}
}

func (python *PythonProject) GenerateMain(rootDir string, appType string) {
	file := "main.py"
	fileManager.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		fileManager.WriteFile(rootDir, file, appImport.ImportForMvcApp())
	case  "api":
		fileManager.WriteFile(rootDir, file, appImport.ImportForRestApi())
	}
	fileManager.WriteFile(rootDir, file, appImport.AppMainCode())
}

func (python *PythonProject) GenerateModels(rootDir string) {
	modelsFolder := rootDir+"/models"
	roleFile := "role.py"
	userFile := "user.py"
	taskFile := "task.py"
	fileManager.CreateFolderAll(modelsFolder)
	fileManager.CreateFile(modelsFolder, roleFile)
	fileManager.CreateFile(modelsFolder, taskFile)
	fileManager.CreateFile(modelsFolder, userFile)
	fileManager.WriteFile(modelsFolder, roleFile, model.RoleModel())
	fileManager.WriteFile(modelsFolder, userFile, model.UserModel())
	fileManager.WriteFile(modelsFolder, taskFile, model.TaskModel())
}

func (python *PythonProject) GenerateViews(rootDir string) {
	fileManager.CreateFolderAll(rootDir+"/templates")
}

func (python *PythonProject) GenerateStaticFiles(rootDir string) {
	staticCss := rootDir+"/static/css"
	staticJs := rootDir+"/static/js"
	staticImgs := rootDir+"/static/images"
	jsFile := "script.js"
	cssFile := "style.css"
	fileManager.CreateFolderAll(staticCss)
	fileManager.CreateFolderAll(staticJs)
	fileManager.CreateFolderAll(staticImgs)
	fileManager.CreateFile(staticCss, cssFile)
	fileManager.CreateFile(staticJs, jsFile)
	fileManager.WriteFile(staticCss, cssFile, staticFile.CssContent())
	fileManager.WriteFile(staticJs, jsFile, staticFile.JsContent())
}

func (python *PythonProject) GenerateMvcControllers(rootDir string) {
	controllersFolder := rootDir+"/controllers"
	roleFile := "role_controller.py"
	userFile := "user_controller.py"
	taskFile := "task_controller.py"
	authFile := "auth_controller.py"
	fileManager.CreateFolderAll(controllersFolder)
	fileManager.CreateFile(controllersFolder, roleFile)
	fileManager.CreateFile(controllersFolder, taskFile)
	fileManager.CreateFile(controllersFolder, userFile)
	fileManager.CreateFile(controllersFolder, authFile)
	fileManager.WriteFile(controllersFolder, roleFile, mvcController.RoleController())
	fileManager.WriteFile(controllersFolder, userFile, mvcController.UserController())
	fileManager.WriteFile(controllersFolder, taskFile, mvcController.TaskController())
	fileManager.WriteFile(controllersFolder, authFile, mvcController.AuthController())
}

func (python *PythonProject) GenerateApiControllers(rootDir string) {
	apiControllersFolder := rootDir+"/api_controllers"
	roleFile := "role_api.py"
	userFile := "user_api.py"
	taskFile := "task_api.py"
	authFile := "auth_api.py"
	fileManager.CreateFolderAll(apiControllersFolder)
	fileManager.CreateFile(apiControllersFolder, roleFile)
	fileManager.CreateFile(apiControllersFolder, taskFile)
	fileManager.CreateFile(apiControllersFolder, userFile)
	fileManager.CreateFile(apiControllersFolder, authFile)
	fileManager.WriteFile(apiControllersFolder, roleFile, apiController.RoleApiController())
	fileManager.WriteFile(apiControllersFolder, userFile, apiController.UserApiController())
	fileManager.WriteFile(apiControllersFolder, taskFile, apiController.TaskApiController())
	fileManager.WriteFile(apiControllersFolder, authFile, apiController.AuthApiController())
}

func (python *PythonProject) GenerateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	fileManager.CreateFolderAll(dbDir)
	fileManager.CreateFile(dbDir, file)
	fileManager.WriteFile(dbDir, file, mysql.GetDatabaseScript())
}

func (python *PythonProject) GeneratePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	fileManager.CreateFolderAll(dbDir)
	fileManager.CreateFile(dbDir, file)
	fileManager.WriteFile(dbDir, file, postgres.GetDatabaseScript())
}

func (python *PythonProject) GenerateRequirements(rootDir string, db string) {
	var deps *scripts.PythonDeps
	file := "requirements.txt"
	fileManager.CreateFile(rootDir, file)
	fileManager.WriteFile(rootDir, file, deps.Requirements(db))
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
	python.GenerateReadme(appName, db, appType)
}