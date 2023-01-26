package pythonsamples

type AppImport struct {}


func (imp* AppImport) ImportForConfig() string  {
	return `
		from flask import Flask
		from flask_sqlalchemy import SQLAlchemy
		from sqlalchemy import create_engine
	`
}

func (imp* AppImport) ImportForAllModels() string  {
	return `
		from config import db, engine
	`
}

func (imp* AppImport) ImportForAllControllers() string  {
	return `
		from config import db, engine
	`
}

func (imp* AppImport) ImportForUserModel() string  {
	return `
		from config import db, engine
		from flask import session
	`
}

func (imp* AppImport) ImportForTaskController() string  {
	return `
		from config import *
		from flask import render_template, request, redirect, url_for
		from models.task import TaskModel
		from models.user import User
	`
}

func (imp* AppImport) ImportForAuthController() string  {
	return `
		from config import *
		from flask import render_template, request, redirect, url_for, session
		from models.user import User
	`
}

func (imp* AppImport) ImportForUserController() string  {
	return `
		import os
		from werkzeug.utils import secure_filename
		from flask import render_template, request, redirect, flash, url_for
		from config import *
		from models.usuario import Usuario
	`
}


func (imp* AppImport) ImportForMvcApp() string  {
	return `
		from flask import jsonify, render_template
		from config import *
		from controllers import (
			user_controller,
			auth_controller,
			task_controller
		)
	`
}


func (imp* AppImport) ImportForRestApi() string  {
	return `
		from flask import jsonify, render_template
		from config import *
		from api_controllers import (
			user_api,
			auth_api,
			task_api
		)
	`
}