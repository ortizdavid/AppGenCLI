package pythonsamples

type MvcController struct {}

var mvcImport *AppImport

func (mvc * MvcController) TaskController()  string {
return ``+mvcImport.ImportForTaskController()+`

class TaskController:
	pass`
}

func (mvc * MvcController) UserController()  string {
return ``+mvcImport.ImportForUserController()+`

class UserController:
	pass`
}

func (mvc * MvcController) RoleController()  string {
return ``+mvcImport.ImportForAllControllers()+`

class RoleController:
	pass`
}

func (mvc * MvcController) AuthController()  string {
return ``+mvcImport.ImportForTaskController()+`

class AuthController:
	pass`
}