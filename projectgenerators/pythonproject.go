package projectgenerators

import (
	"fmt"

	"github.com/ortizdavid/appgen/filemanager"
	"github.com/ortizdavid/appgen/helpers"
	dbsamples "github.com/ortizdavid/appgen/samples/db"
	pythonsamples "github.com/ortizdavid/appgen/samples/python"
	"github.com/ortizdavid/appgen/samples/scripts"
)

type PythonProject struct {}


var pyFileManager *filemanager.FileManager
var pyImport *pythonsamples.AppImport


func (python *PythonProject) GetProjectTypes() []string {
	types := [] string {"api", "mvc"}
	return types
}


func (python *PythonProject) CreateApp(appName string, appType string, db string) {

	if helpers.Contains(python.GetProjectTypes(), appType) == false {
		fmt.Printf(helpers.PROJECT_ERROR)
		fmt.Printf(helpers.UNSUPORTED_TYPE, appType, "Python")
		helpers.ListLanguages()

	} else {
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
		python.GenerateModels(appName)
		python.GenerateConfig(appName, db)
		python.GenerateRequirements(appName, db)
		python.GenerateMain(appName, appType)
		python.GenerateReadme(appName, db, appType)
		python.GenerateGitIgnore(appName, appType)
		pyFileManager.CreateFolderAll(appName+"/uploads/imgs")
		pyFileManager.CreateFolderAll(appName+"/uploads/docs")
	}
}

func (python *PythonProject) GenerateConfig(rootDir string, db string) {
	var config *pythonsamples.ConfigPy
	file := "config.py"
	pyFileManager.CreateFile(rootDir, file)
	pyFileManager.WriteFile(rootDir, file, pyImport.ImportForConfig())
	pyFileManager.WriteFile(rootDir, file, config.CreateConfig(db))
}


func (python *PythonProject) GenerateReadme(rootDir string, db string, appType string) {
	var readme *helpers.ReadMePy
	file := "README.md"
	pyFileManager.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, readme.ReadmeMVC(db))
	case "api":
		pyFileManager.WriteFile(rootDir, file, readme.ReadmeAPI(db))
	}
}


func (python *PythonProject) GenerateGitIgnore(rootDir string, appType string) {
	var ignore *helpers.GitIgnorePy
	file := ".gitignore"
	pyFileManager.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, ignore.GitIgnoreMvc())
	case "api":
		pyFileManager.WriteFile(rootDir, file, ignore.GitIgnoreAPI())
	}
}


func (python *PythonProject) GenerateMain(rootDir string, appType string) {
	file := "main.py"
	pyFileManager.CreateFile(rootDir, file)
	switch appType {
	case "mvc":
		pyFileManager.WriteFile(rootDir, file, pyImport.ImportForMvcApp())
	case  "api":
		pyFileManager.WriteFile(rootDir, file, pyImport.ImportForRestApi())
	}
	pyFileManager.WriteFile(rootDir, file, pyImport.AppMainCode())
}


func (python *PythonProject) GenerateMySqlDB(rootDir string) {
	var mysql *dbsamples.MySqlDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	pyFileManager.CreateFolderAll(dbDir)
	pyFileManager.CreateFile(dbDir, file)
	pyFileManager.WriteFile(dbDir, file, mysql.GetDatabaseScript())
}


func (python *PythonProject) GeneratePostgresDB(rootDir string) {
	var postgres *dbsamples.PostgresDB
	dbDir := rootDir+"/database"
	file := "db_task.sql"
	pyFileManager.CreateFolderAll(dbDir)
	pyFileManager.CreateFile(dbDir, file)
	pyFileManager.WriteFile(dbDir, file, postgres.GetDatabaseScript())
}


func (python *PythonProject) GenerateRequirements(rootDir string, db string) {
	var deps *scripts.PythonDeps
	file := "requirements.txt"
	pyFileManager.CreateFile(rootDir, file)
	pyFileManager.WriteFile(rootDir, file, deps.Requirements(db))
}


func (python *PythonProject) GenerateHelpers(rootDir string) {
	var helper *pythonsamples.Helper
	helpersFolder := rootDir+"/helpers"
	httpFile := "http_code.py"
	constFile := "constants.py"
	pyFileManager.CreateFolderAll(helpersFolder)
	pyFileManager.CreateFile(helpersFolder, httpFile)
	pyFileManager.CreateFile(helpersFolder, constFile)
	pyFileManager.WriteFile(helpersFolder, httpFile, helper.HttpCodes())
	pyFileManager.WriteFile(helpersFolder, constFile, helper.Constants())
}


func (python *PythonProject) GenerateModels(rootDir string) {
	var model *pythonsamples.Model
	modelsFolder := rootDir+"/models"
	roleFile := "role.py"
	userFile := "user.py"
	taskFile := "task.py"
	pyFileManager.CreateFolderAll(modelsFolder)
	pyFileManager.CreateFile(modelsFolder, roleFile)
	pyFileManager.CreateFile(modelsFolder, taskFile)
	pyFileManager.CreateFile(modelsFolder, userFile)
	pyFileManager.WriteFile(modelsFolder, roleFile, model.RoleModel())
	pyFileManager.WriteFile(modelsFolder, userFile, model.UserModel())
	pyFileManager.WriteFile(modelsFolder, taskFile, model.TaskModel())
}


func (python *PythonProject) GenerateStaticFiles(rootDir string) {
	var staticFile *pythonsamples.StaticFile
	staticCss := rootDir+"/static/css"
	staticJs := rootDir+"/static/js"
	staticImgs := rootDir+"/static/images"
	jsFile := "script.js"
	cssFile := "style.css"
	pyFileManager.CreateFolderAll(staticCss)
	pyFileManager.CreateFolderAll(staticJs)
	pyFileManager.CreateFolderAll(staticImgs)
	pyFileManager.CreateFile(staticCss, cssFile)
	pyFileManager.CreateFile(staticJs, jsFile)
	pyFileManager.WriteFile(staticCss, cssFile, staticFile.CssContent())
	pyFileManager.WriteFile(staticJs, jsFile, staticFile.JsContent())
}


func (python *PythonProject) GenerateMvcControllers(rootDir string) {
	var mvcController *pythonsamples.MvcController
	controllersFolder := rootDir+"/controllers"
	roleFile := "role_controller.py"
	userFile := "user_controller.py"
	taskFile := "task_controller.py"
	authFile := "auth_controller.py"
	pyFileManager.CreateFolderAll(controllersFolder)
	pyFileManager.CreateFile(controllersFolder, roleFile)
	pyFileManager.CreateFile(controllersFolder, taskFile)
	pyFileManager.CreateFile(controllersFolder, userFile)
	pyFileManager.CreateFile(controllersFolder, authFile)
	pyFileManager.WriteFile(controllersFolder, roleFile, mvcController.RoleController())
	pyFileManager.WriteFile(controllersFolder, userFile, mvcController.UserController())
	pyFileManager.WriteFile(controllersFolder, taskFile, mvcController.TaskController())
	pyFileManager.WriteFile(controllersFolder, authFile, mvcController.AuthController())
}


func (python *PythonProject) GenerateApiControllers(rootDir string) {
	var apiController *pythonsamples.ApiController
	apiControllersFolder := rootDir+"/api_controllers"
	roleFile := "role_api.py"
	userFile := "user_api.py"
	taskFile := "task_api.py"
	authFile := "auth_api.py"
	pyFileManager.CreateFolderAll(apiControllersFolder)
	pyFileManager.CreateFile(apiControllersFolder, roleFile)
	pyFileManager.CreateFile(apiControllersFolder, taskFile)
	pyFileManager.CreateFile(apiControllersFolder, userFile)
	pyFileManager.CreateFile(apiControllersFolder, authFile)
	pyFileManager.WriteFile(apiControllersFolder, roleFile, apiController.RoleApiController())
	pyFileManager.WriteFile(apiControllersFolder, userFile, apiController.UserApiController())
	pyFileManager.WriteFile(apiControllersFolder, taskFile, apiController.TaskApiController())
	pyFileManager.WriteFile(apiControllersFolder, authFile, apiController.AuthApiController())
}


func (python *PythonProject) GenerateViews(rootDir string) {

	var layout *pythonsamples.Layout
	var perr *pythonsamples.PageError
	var auth *pythonsamples.AuthTemplate
	var user *pythonsamples.UserTemplate
	var task *pythonsamples.TaskTemplate
	var role *pythonsamples.RoleTemplate

	templatesFolder := rootDir+"/templates"
	layoutsFolder := templatesFolder+"/layouts"
	errorFolder := templatesFolder+"/error"
	authFolder := templatesFolder+"/auth"
	userFolder := templatesFolder+"/user"
	taskFolder := templatesFolder+"/task"
	roleFolder := templatesFolder+"/role"
	
	backLayoutFile := "front-layout.html"
	frontLayoutFile := "back-layout.html"
	normalMenuFile := "normal-menu.html"
	adminMenuFile := "admin-menu.html"

	loginFile := "login.html"
	homeFile := "home.html"
	err404File := "404.html"
	
	userAddFile := "add.html"
	userEditFile := "edit.html"
	userDataFile := "user-data.html"
	userSearchFile := "search.html"
	userShowFile := "show.html"
	userDetailsFile := "details.html"

	taskAddFile := "add.html"
	taskEditFile := "edit.html"
	taskShowFile := "show.html"
	taskSearchFile := "search.html"
	taskDetailsFile := "details.html"

	roleAddFile := "add.html"
	roleShowFile := "show.html"
	roleDetailsFile := "details.html"

	pyFileManager.CreateFolderAll(templatesFolder)
	pyFileManager.CreateFolderAll(layoutsFolder)
	pyFileManager.CreateFolderAll(authFolder)
	pyFileManager.CreateFolderAll(userFolder)
	pyFileManager.CreateFolderAll(roleFolder)
	pyFileManager.CreateFolderAll(taskFolder)
	pyFileManager.CreateFolderAll(errorFolder)
	
	pyFileManager.CreateFile(errorFolder, err404File)
	pyFileManager.CreateFileAll(authFolder, loginFile, homeFile)
	pyFileManager.CreateFileAll(layoutsFolder, frontLayoutFile, backLayoutFile, normalMenuFile, adminMenuFile)
	pyFileManager.CreateFileAll(roleFolder, roleAddFile, roleShowFile, roleDetailsFile)
	pyFileManager.CreateFileAll(taskFolder, taskAddFile, taskEditFile, taskSearchFile, taskShowFile, taskDetailsFile)
	pyFileManager.CreateFileAll(userFolder, userAddFile, userEditFile, userDataFile, userSearchFile, userShowFile, userSearchFile, userDetailsFile)

	pyFileManager.WriteFile(errorFolder, err404File, perr.Error404())
	pyFileManager.WriteFile(layoutsFolder, frontLayoutFile, layout.FontLayout())
	pyFileManager.WriteFile(layoutsFolder, backLayoutFile, layout.BackLayout())
	pyFileManager.WriteFile(layoutsFolder, adminMenuFile, layout.AdminMenu())
	pyFileManager.WriteFile(layoutsFolder, normalMenuFile, layout.NormalMenu())

	pyFileManager.WriteFile(authFolder, loginFile, auth.LoginTemplate())
	pyFileManager.WriteFile(authFolder, homeFile, auth.HomeTemplate())

	pyFileManager.WriteFile(userFolder, userAddFile, user.AddTemplate())
	pyFileManager.WriteFile(userFolder, userEditFile, user.EditTemplate())
	pyFileManager.WriteFile(userFolder, userShowFile, user.ShowTemplate())
	pyFileManager.WriteFile(userFolder, userDetailsFile, user.DetailsTemplate())
	pyFileManager.WriteFile(userFolder, userSearchFile, user.SearchTemplate())
	pyFileManager.WriteFile(userFolder, userDataFile, user.UserDataTemplate())

	pyFileManager.WriteFile(taskFolder, taskAddFile, task.AddTemplate())
	pyFileManager.WriteFile(taskFolder, taskEditFile, task.EditTemplate())
	pyFileManager.WriteFile(taskFolder, taskShowFile, task.ShowTemplate())
	pyFileManager.WriteFile(taskFolder, taskSearchFile, task.SearchTemplate())
	pyFileManager.WriteFile(taskFolder, taskDetailsFile, task.DetailsTemplate())

	pyFileManager.WriteFile(roleFolder, roleAddFile, role.AddTemplate())
	pyFileManager.WriteFile(taskFolder, roleDetailsFile, role.DetailsTemplate())
	pyFileManager.WriteFile(taskFolder, roleShowFile, role.ShowTemplate())
}