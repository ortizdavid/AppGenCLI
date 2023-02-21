package pythonsamples

type AppImport struct {}


func (imp *AppImport) ParticularImport(appType string) string {
	var strImport string
	switch appType {
	case "mvc":
		strImport = "from flask import render_template, request, redirect, url_for, session"
	case "api":
		strImport = "from helpers.http_code import *\nfrom flask import jsonify, request, session"
	}
	return strImport
}


func (imp* AppImport) ImportForConfig() string  {
return `from flask import Flask
from sqlalchemy import create_engine
from flask_sqlalchemy import SQLAlchemy`
}


func (imp* AppImport) ImportForAllModels() string  {
return `from config import db, engine
from sqlalchemy import text`
}


func (imp* AppImport) ImportForUserModel() string  {
return `from config import *
from sqlalchemy import text
from flask import session
from config import db, engine`
}



func (imp* AppImport) ImportForRoleController(appType string) string  {
return `from config import *
from models.role import Role
from models.user import User
`+imp.ParticularImport(appType)+``
}
	


func (imp* AppImport) ImportForTaskController(appType string) string  {
return `from config import *
from models.user import User
from models.task import Task
`+imp.ParticularImport(appType)+``
}


func (imp* AppImport) ImportForAuthController(appType string) string  {
return `from config import *
from helpers.password_handler import *
from models.user import User
`+imp.ParticularImport(appType)+``
}


func (imp* AppImport) ImportForUserController(appType string) string  {
return `from config import *
from helpers.file_uploader import *
from helpers.password_handler import *
from models.user import User
from models.role import Role
`+imp.ParticularImport(appType)+``
}


func (imp* AppImport) ImportForMvcApp() string  {
return `from config import *
from flask import jsonify, render_template
from controllers import (
	user_controller,
	auth_controller,
	task_controller,
	front_controller
)
`
}


func (imp* AppImport) ImportForRestApi() string  {
return `from config import *
from api_controllers import (
	user_api,
	auth_api,
	task_api,
	role_api,
	root_api
)
`
}


func (imp* AppImport) AppMainCode() string {
return `
if __name__ == '__main__':
    app.run(port=APP_PORT, debug=True)`
}