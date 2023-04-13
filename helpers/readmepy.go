package helpers

type ReadMePy struct {}

func (rdm *ReadMePy) ReadmeAPI(db string)  string {
return `# Python REST API with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}


func (rdm *ReadMePy) ReadmeMVC(db string)  string {
return `# Python MVC APP with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}


func (rdm *ReadMePy) Instrunctions(db string) string {
return`## Pre requisites:
- `+StrDatabase(db)+`
- Python: version 3.9 
- Pip
- Python Virtual environment (venv)

## Steps for run application: 
- Go to project directory: cd ProjectName
- Copy database script [db_task.sql](database/db_task.sql) to database management system
- Install virtual environment: pip install virtualenv
- Create virtualenvironment: virtualenv venv
- Activate virtual environment: venv/Scripts/activate
- Install all dependencies in [requirements.txt](requirements.txt): pip install -r requirements.txt
- Change database connection variables on: [config.py](config.py)
- Run application: flask run or python app.py
- Open browser: http://localhost:5000
- Users for tests: admin01, admin02, user1, user2, 
- Passwords for all users: 12345678`
}

func (rdm *ReadMePy) InstrunctionsBeforeRun(appName string) string {
return `
Instructions for run `+appName+`:
- Go project directory: cd `+appName+`
- Copy database script 'database/db_task.sql' to database management system
- Install virtual environment: pip install virtualenv
- Create virtualenvironment: virtualenv venv
- Activate virtual environment: venv/Scripts/activate
- Install all dependencies in 'requirements.txt': pip install -r requirements.txt
- Change database connection variables on: 'config.py'
- Run application: flask run or python app.py
- Open browser: http://localhost:5000
- Users for tests: admin01, admin02, user1, user2, ...
- Passwords for all users: 12345678
`	
}