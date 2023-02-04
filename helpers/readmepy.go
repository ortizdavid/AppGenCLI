package helpers

type ReadMePy struct {}

func (rdm *ReadMePy) ReadmeAPI(db string)  string {
return `
# Python REST API with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}

func (rdm *ReadMePy) ReadmeMVC(db string)  string {
return `
# Python MVC APP with `+StrDatabase(db)+`
This API manage user tasks
Users can Add, Remove, List and Edit daily tasks
`+rdm.Instrunctions(db)+``
}

func (rdm *ReadMePy) Instrunctions(db string) string {
return`
## Pre requisites:
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