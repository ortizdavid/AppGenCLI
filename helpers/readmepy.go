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
- Copy database script: [db_task.sql](database/db_task.sql)
- Install virtual environment: pip install virtualenv
- Create virtualenvironment: virtualenv venv
- Activate virtual environment: venv/Scripts/activate
- Install all dependencies in [requirements.txt](requirements.txt): pip install -r requirements.txt
- Configure database on: [config.py](config.py)
- Run application: flask run or python app.py
- Access the application with URL: http://localhost:5000
- Users for tests: admin01, admin02, user1, user2, ...
- Passwords for all users: 12345678`
}

func (rdm *ReadMePy) InstrunctionsBeforeRun(appName string) string {
return `
Instructions for run `+appName+`:
1. Go project directory: cd `+appName+`
2. Copy script ' database/db_task.sql ' to database management system
3. Install all Virtual Environment: pip install virtualenv
4. Create virtualenv: virtualenv venv
5. Activate virtualenv: venv/Scripts/activate
6. Install all dependencies: pip install -r requirements.txt
7. Configure database DB_NAME, DB_PASSWORD: config.py
8. Run application: flask run or python app.py
9. Open browser: http://localhost:5000
10. Users for tests: admin01, admin02, user1, user2, ...
11. Passwords for all users: 12345678
`	
}