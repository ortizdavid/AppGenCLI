package pythonsamples

type ApiController struct {}


var apiImport *AppImport


func (api *ApiController) TaskApiController()  string {
return ``+apiImport.ImportForTaskController("api")+`

class TaskApi:
	
	@app.route(f'/{API_ROOT}/tasks', methods=['GET'])
	def get_all_tasks():
		tasks = [task.to_json() for task in Task.get_all()]
		num_rows = len(tasks)
		if num_rows == 0:
			return jsonify({'tasks': tasks, 'num_rows': num_rows}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'tasks': tasks, 'num_rows': num_rows}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/tasks/<id>', methods=['GET'])
	def get_task(id):
		task = Task.get_by_id(id)
		if task:
			return jsonify(task.to_json()), HTTP_CODE_OK
		else:
			return jsonify({'message': 'Task not found'}), HTTP_CODE_NOT_FOUND`
}


func (api *ApiController) UserApiController()  string {
return ``+apiImport.ImportForUserController("api")+`

class UserApi:
	pass`
}


func (api *ApiController) RoleApiController()  string {
return ``+apiImport.ImportForRoleController("api")+`

class RoleApi:
	pass`
}


func (api *ApiController) AuthApiController()  string {
return ``+apiImport.ImportForAuthController("api")+`

class AuthApi:
	pass`
}