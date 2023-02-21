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
			return jsonify({'message': 'Error while adding Task', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
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
			return jsonify({'message': 'Error while updating Task', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': task.to_json(), 'message': 'Task Updated', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/tasks/<id>', methods=['DELETE'])
	def delete_task(id):
		task = Task.get_by_id(id)
		if task is None:
			return jsonify({'message': 'Error While Deleting Task', 'success': False}), HTTP_CODE_NOT_FOUND
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

	@app.route(f'/{API_ROOT}/users', methods=['GET'])
	def get_all_users():
		users = [user.to_json() for user in User.get_all()]
		num_rows = len(users)
		if num_rows == 0:
			return jsonify({'users': users, 'num_rows': num_rows, 'message': 'No Results Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'users': users, 'num_rows': num_rows, 'message': 'Users Found', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/users/<id>', methods=['GET'])
	def get_user(id):
		user = User.get_by_id(id)
		if user is None:
			return jsonify({'message': 'User Not found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			return jsonify({'data': user.to_json(), 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/users', methods=['POST'])
	def add_user():
		data = request.get_json()
		role_id = data['role_id']
		user_name = data['user_name']
		password = data['password']
		image = ""

		if User.exists(user_name):
			return jsonify({'message': f"User '{user_name}' already exists", 'sucess': False}), HTTP_CODE_BAD_REQUEST
		try:
			user = User(role_id, user_name, password, image)
			user.save()
		except:
			return jsonify({'message': 'Error while adding user', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': user.to_json(), 'message': 'User Added', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/users/<id>', methods=['PUT'])
	def edit_user(id):
		user = User.get_by_id(id)
		data = request.get_json()
		data = request.get_json()
		role_id = data['role_id']
		user_name = data['user_name']
		password = data['password']
		image = ""
		try:
			if user is None:
				user = User(role_id, user_name, password, image)
			else:
				user.role_id = role_id
				user.user_name = user_name
				user.password = password
				user.image = image
			user.save()
		except:
			return jsonify({'message': 'Error while updating User', 'success': False}) , HTTP_CODE_INTERNAL_ERROR
		return jsonify({'data': user.to_json(), 'message': 'User Updated', 'success': True}), HTTP_CODE_CREATED


	@app.route(f'/{API_ROOT}/users/<id>', methods=['DELETE'])
	def delete_user(id):
		user = User.get_by_id(id)
		if user is None:
			return jsonify({'message': 'Error while deleting User', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			user.delete()
			return jsonify({'message': 'User Deleted!', 'success': True}), HTTP_CODE_OK


	@app.route(f'/{API_ROOT}/users/<id>/upload', methods=['POST'])
	def upload_image(id):
		user = User.get_by_id(id)
		if user is None:
			return jsonify({'message' : 'User not Found', 'success': False}), HTTP_CODE_NOT_FOUND
		else:
			uploader = FileUploader()
			image = uploader.upload_image('image', UPLOAD_DIR_IMGS)
			errors = uploader.errors
			if len(errors) > 0:
				return jsonify({'errors': errors, 'message' : 'Error while Uploading', 'success': False}), HTTP_CODE_BAD_REQUEST
			else:
				user.update_image(image, id)
				return jsonify({'file': image, 'message' : 'File successfully uploaded', 'success': True}), HTTP_CODE_CREATED`
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

	@app.route(f'/{API_ROOT}/auth/login', methods=['POST'])
	def login():
		data = request.get_json()
		user_name = data['user_name']
		password = data['password']
		user = User.get_by_username(user_name)
		encrypted_password = user.password
		
		if PasswordHandler.check(encrypted_password, password):
			session['user_name'] = user_name
			session['password'] = encrypted_password
			logged_user = User.get_logged_user_basic()
			return jsonify({'data': logged_user.to_json(), 'message': 'Login Succeeded', 'success': True}), HTTP_CODE_OK
		else:
			return jsonify({'message': 'Invalid Username or Password', 'success': False}), HTTP_CODE_NOT_FOUND


	@app.route(f'/{API_ROOT}/auth/logout', methods=['GET'])
	def logout():
		if 'user_name' in session:
			session.pop('user_name')
			return jsonify({'message': 'Logout Succeeded', 'success': True}), HTTP_CODE_OK
		else:
			return jsonify({'message': 'Logout Failed', 'success': False}), HTTP_CODE_BAD_REQUEST


	@app.route(f'/{API_ROOT}/auth/user', methods=['GET'])
	def get_current_user():
		if 'user_name' in session:
			logged_user = User.get_logged_user_basic()
			return jsonify({'data': logged_user.to_json(), 'message': 'Logged User Found', 'success': True}), HTTP_CODE_OK
		else:
			return jsonify({'message': 'Logged User Not Found', 'success': False}), HTTP_CODE_NOT_FOUND`
}


func (api *ApiController) RootApiController()  string {
return `from flask import send_file
from config import app

class RootAPI:

    @app.route('/', methods=['GET'])
    @app.route('/api', methods=['GET'])
    def root():
        return  """
            <h1>REST API for Task Management</h1>
            <p>Basic End Points:</p>
            <ul>
                <li>Users: /api/tasks</li>
                <li>Tasks: /api/tasks</li>
                <li>Roles: /api/roles</li>
                <li>Login: /api/auth/login</li>
                <li>Logout: /api/auth/logout</li>
            </ul>
            <p>Postman Collections:</p>
            <ul>
                <li>Postman Collection <a href="/postman-collection">Download</a></li>
                <li>Postman Test Run <a href="/postman-test">Download</a></li>
            </ul>
        """

    @app.route('/postman-collection', methods=['GET'])
    def postman_collection():
        return send_file("__postman_collections__/postman_collection.json", as_attachment='postman-collection.json')

    @app.route('/postman-test', methods=['GET'])
    def postman_test():
        return send_file("__postman_collections__/postman_test_run.json", as_attachment='postman-test.json')`
}