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
Instructions:
1. Go project directory:
	cd `+appName+`
2. Copy database script to database management system:
	database/db_task.sql 
3. Install all dependencies:
	pip install -r requirements.txt
4. Configure database DB_NAME, DB_PASSWORD
	config.py
5. Run application: 
	flask run
	python app.py
6. Open browser:
	http://localhost:5000
`	
}