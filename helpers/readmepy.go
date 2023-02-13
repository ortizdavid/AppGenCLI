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
- Install all dependencies in [requirements.txt](requirements.txt): pip install -r requirements.txt
- Configure database on: [config.py](config.py)
- Run application: flask run or python app.py
- Open browser with URL: http://localhost:5000`
}

func (rdm *ReadMePy) InstrunctionsBeforeRun(appName string) string {
return `
Instructions for run `+appName+`:
1. Go project directory: cd `+appName+`
2. Copy script ' database/db_task.sql ' to database management system
3. Install all Virtual Environment: pip install virtualenv
4. Activate virtualenv: venv/Scripts/activate
5. Install all dependencies: pip install -r requirements.txt
6. Configure database DB_NAME, DB_PASSWORD: config.py
7. Run application: flask run or python app.py
8. Open browser: http://localhost:5000
`	
}