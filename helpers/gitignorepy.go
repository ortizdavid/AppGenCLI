package helpers

type GitIgnorePy struct {}

func (ign *GitIgnorePy) GitIgnoreAPI()  string {
return `venv/

*.pyc
__pycache__/
`
}

func (ign *GitIgnorePy) GitIgnoreMvc()  string {
return `venv/
uploads/

*.pyc
__pycache__/
`
}

