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
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'No Results Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'Tasks Found', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/tasks/<id>', methods=['GET'])
	def get_task(id):
		task = Task.get_by_id(id)
		if task is None:
			return jsonify({'message': 'Task Not found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'data': task.to_json(), 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/tasks', methods=['POST'])
	def add_task():
		data = request.get_json()
		task_name = data['task_name']
		start_date = data['start_date']
		end_date = data['end_date']
		description = data['description']
		user_id = data['user_id']

		if Task.exists(user_id, task_name, start_date, end_date):
			return jsonify({'message': 'Task already exists', 'sucess': False}), HTTP_CODE_BAD_REQUEST
		try:
			task = Task(user_id, task_name, start_date, end_date, description)
			task.save()
		except:
			return jsonify({'message': 'Erro while adding Task', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': task.to_json(), 'message': 'Task Added', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/tasks/<id>', methods=['PUT'])
	def edit_task(id):
		task = Task.get_by_id(id)
		data = request.get_json()
		task_name = data['task_name']
		start_date = data['start_date']
		end_date = data['end_date']
		description = data['description']
		user_id = data['user_id']
		try:
			if task is None:
				task = Task(user_id, task_name, start_date, end_date, description)
			else:
				task.task_name = task_name
				task.start_date = start_date
				task.end_date = end_date
				task.description = description
				task.user_id = user_id
			task.save()
		except:
			return jsonify({'message': 'Erro while updating Task', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': task.to_json(), 'message': 'Task Updated', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/tasks/<id>', methods=['DELETE'])
	def delete_task(id):
		task = Task.get_by_id(id)
		if task is None:
			return jsonify({'message': 'Task Not Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			task.delete()
			return jsonify({'message': 'Task Deleted!', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/tasks/<start_date>/<end_date>', methods=['GET'])
	def filter_task_by_date(start_date, end_date):
		tasks = [task.to_json() for task in Task.get_by_date(start_date, end_date)]
		num_rows = len(tasks)
		if num_rows == 0:
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'No Results Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'Tasks Found', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/tasks/user/<id>', methods=['GET'])
	def filter_task_by_user_id(id):
		tasks = [task.to_json() for task in Task.get_by_user(id)]
		num_rows = len(tasks)
		if num_rows == 0:
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'No Results Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'tasks': tasks, 'num_rows': num_rows, 'message': 'Tasks Found', 'success': True}), HTTP_CODE_OK`
}


func (api *ApiController) UserApiController()  string {
return ``+apiImport.ImportForUserController("api")+`

class UserApi:
	pass`
}


func (api *ApiController) RoleApiController()  string {
return ``+apiImport.ImportForRoleController("api")+`

class RoleApi:
	
	@app.route(f'/{API_ROOT}/roles', methods=['GET'])
	def get_all_roles():
		roles = [role.to_json() for role in Role.get_all()]
		num_rows = len(roles)
		if num_rows == 0:
			return jsonify({'roles': roles, 'num_rows': num_rows, 'message': 'No Results Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'roles': roles, 'num_rows': num_rows, 'message': 'Roles Found', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/roles/<id>', methods=['GET'])
	def get_role(id):
		role = Role.get_by_id(id)
		if role is None:
			return jsonify({'message': 'Role Not found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'data': role.to_json(), 'message': 'Role Found', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/roles', methods=['POST'])
	def add_role():
		data = request.get_json()
		role_name = data['role_name']
		if Role.exists(role_name):
			return jsonify({'message': 'Role already exists', 'sucess': False}), HTTP_CODE_BAD_REQUEST
		try:
			role = Role(role_name)
			role.save()
		except:
			return jsonify({'message': 'Erro while adding Role', 'success': False}), HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': role.to_json(), 'message': 'Role Added', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/roles/<id>', methods=['PUT'])
	def edit_role(id):
		role = Role.get_by_id(id)
		data = request.get_json()
		role_name = data['role_name']
		try:
			if role is None:
				role = Role(role_name)
			else:
				role.role_name = role_name
			role.save()
		except:
			return jsonify({'message': 'Erro while updating Role', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': role.to_json(), 'message': 'Role Updated', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/roles/<id>', methods=['DELETE'])
	def delete_role(id):
		role = Role.get_by_id(id)
		if role is None:
			return jsonify({'message': 'role Not Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			role.delete()
			return jsonify({'message': 'role Deleted!', 'success': True}), HTTP_CODE_OK`
}


func (api *ApiController) AuthApiController()  string {
return ``+apiImport.ImportForAuthController("api")+`

class AuthApi:
	pass`
}