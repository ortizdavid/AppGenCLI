package scripts

type PythonDeps struct {}


func (pyd * PythonDeps) InstallVenv() string {
return `
pip install virtualenv
`
}


func (pyd * PythonDeps) ActivateVenv() string {
return `
venv/Scripts/activate
`
}


func (pyd * PythonDeps) InstallDeps() string {
return `
pip install Flask
pip install SQLAlchemy
pip install flask-SQLAlchemy
pip install psycopg2
pip install psycopg2
pip install Werkzeug
`
}


func (pyd * PythonDeps) Requirements(db string) string {
	var dbDep string
	switch db {
	case "mysql":
		dbDep = "MySQL\npymysql"
	case "postgres":	
		dbDep = "psycopg2"
	}
	return `werkzeug
Flask
flask-SQLAlchemy
SQLAlchemy
`+dbDep+`
`
}