package pythonsamples

type ApiController struct {}


var apiImport *AppImport


func (api *ApiController) TaskApiController()  string {
return ``+apiImport.ImportForTaskController()+`

class TaskApi:
	pass`
}


func (api *ApiController) UserApiController()  string {
return ``+apiImport.ImportForUserController()+`

class UserApi:
	pass`
}


func (api *ApiController) RoleApiController()  string {
return ``+apiImport.ImportForRoleController()+`

class RoleApi:
	pass`
}


func (api *ApiController) AuthApiController()  string {
return ``+apiImport.ImportForAuthController()+`

class AuthApi:
	pass`
}