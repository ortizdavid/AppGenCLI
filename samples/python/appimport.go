package pythonsamples

type AppImport struct {}

func (imp* AppImport) ImportForConfig() string  {
return `from flask import Flask
from sqlalchemy import create_engine
from flask_sqlalchemy import SQLAlchemy`
}

func (imp* AppImport) ImportForAllModels() string  {
return `from config import db, engine`
}

func (imp* AppImport) ImportForRoleController() string  {
return `from config import *
from models.role import Role
from models.user import User
from flask import render_template, request, redirect, url_for`
}

func (imp* AppImport) ImportForUserModel() string  {
return `from config import *
from flask import session
from config import db, engine`
}

func (imp* AppImport) ImportForTaskController() string  {
return `from config import *
from models.user import User
from models.task import TaskModel
from flask import render_template, request, redirect, url_for, session`
}


func (imp* AppImport) ImportForAuthController() string  {
return `from config import *
from models.user import User
from flask import render_template, request, redirect, url_for, session`
}


func (imp* AppImport) ImportForUserController() string  {
return `import os
from config import *
from models.user import User
from werkzeug.utils import secure_filename
from flask import render_template, request, redirect, flash, url_for`
}


func (imp* AppImport) ImportForMvcApp() string  {
return `from config import *
from flask import jsonify, render_template
from controllers import (
	user_controller,
	auth_controller,
	task_controller
)
`
}

func (imp* AppImport) ImportForRestApi() string  {
return `from flask import jsonify, render_template
from config import *
from api_controllers import (
	user_api,
	auth_api,
	task_api
)
`
}

func (imp* AppImport) AppMainCode() string {
return `
if __name__ == '__main__':
    app.run(port=APP_PORT, debug=True)`
}