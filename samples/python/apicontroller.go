package pythonsamples

type ApiController struct {}

var apiImport *AppImport

func (api *ApiController) TaskApiController()  string {
return ``+apiImport.ImportForTaskController()+`

class TaskController:
	pass`
}

func (api *ApiController) UserApiController()  string {
return ``+apiImport.ImportForUserController()+`

class TaskController:
	pass`
}

func (api *ApiController) RoleApiController()  string {
return ``+apiImport.ImportForAllControllers()+`

class TaskController:
	pass`
}

func (api *ApiController) AuthApiController()  string {
return ``+apiImport.ImportForAuthController()+`

class TaskController:
	pass`
}