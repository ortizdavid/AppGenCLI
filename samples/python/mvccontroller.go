package pythonsamples

type MvcController struct {}


var mvcImport *AppImport


func (mvc * MvcController) TaskController()  string {
return ``+mvcImport.ImportForTaskController("mvc")+`

class TaskController:


	@app.route('/tasks', methods=['GET'])
	def show_tasks():
		tasks = Task.get_all_data()
		num_rows = len(tasks)
		return render_template('task/show.html', tasks=tasks, num_rows=num_rows, logged_user=User.get_logged_user())


	@app.route('/my-tasks', methods=['GET'])
	def user_tasks():
		logged_user=User.get_logged_user()
		tasks = Task.get_all_user_tasks(logged_user.user_id)
		num_rows = len(tasks)
		return render_template('task/user-tasks.html', tasks=tasks, num_rows=num_rows, logged_user=logged_user)


	@app.route('/tasks/<id>/details', methods=['GET'])
	def task_details(id):
		task = Task.get_data_by_id(id)
		if task:
			return render_template('task/details.html', task=task, logged_user=User.get_logged_user())
		else:
			return render_template('errorr/404.html')


	@app.route('/tasks/add', methods=['GET', 'POST'])
	def add_task():
		logged_user=User.get_logged_user()
		if request.method == 'GET': 
			return render_template('task/add.html', logged_user=logged_user)
		else:
			task_name = request.form['task_name']
			start_date = request.form['start_date']
			end_date = request.form['end_date']
			description = request.form['description']
			user_id = logged_user.user_id
			task = Task(user_id, task_name, start_date, end_date, description)
			task.save()
			return redirect(url_for('user_tasks'))


	@app.route('/tasks/<id>/edit', methods=['GET', 'POST'])
	def edit_task(id):
		task = Task.get_by_id(id)
		logged_user=User.get_logged_user()
		if request.method == 'GET': 
			return render_template('task/edit.html', task=task, logged_user=logged_user)
		else:
			task_name = request.form['task_name']
			start_date = request.form['start_date']
			end_date = request.form['end_date']
			description = request.form['description']
			user_id = logged_user.user_id
			new_task = Task(user_id, task_name, start_date, end_date, description)
			new_task.save()
			return redirect(url_for('show_users'))


	@app.route('/task/search', methods=['GET', 'POST'])
	def search_task():
		if request.method == 'GET': 
			return render_template('task/search.html',logged_user=User.get_logged_user())
		else:
			value = request.form['search_value']
			res = Task.search(value)
			num_rows =  len(res)
			return render_template('task/search-results.html', value=value, results=res, 
					num_rows=num_rows, logged_user=User.get_logged_user())`
}


func (mvc * MvcController) UserController()  string {
return ``+mvcImport.ImportForUserController("mvc")+`

class UserController:

	@app.route('/users', methods=['GET'])
	def show_users():
		users = User.get_all_data()
		num_rows = len(users)
		return render_template('user/show.html', users=users, num_rows=num_rows, logged_user=User.get_logged_user())


	@app.route('/users/<id>/details', methods=['GET'])
	def user_details(id):
		user = User.get_data_by_id(id)
		if user:
			return render_template('user/details.html', user=user, logged_user=User.get_logged_user())
		else:
			return render_template('errorr/404.html')


	@app.route('/user-data', methods=['GET'])
	def get_user_data():
		logged_user = User.get_logged_user()
		data = User.get_data_by_id(logged_user.user_id)
		return render_template('user/user-data.html', data=data, logged_user=logged_user)


	@app.route('/users/add', methods=['GET', 'POST'])
	def add_user():
		roles = Role.get_all()
		logged_user=User.get_logged_user()

		if request.method == 'GET': 
			return render_template('user/add.html', logged_user=logged_user, roles=roles)
		else:
			user_name = request.form['user_name']
			role_id = request.form['role_id']
			password = request.form['password']
			encrypted_password = PasswordHandler.generate(password)
			image = ''
			user = User(role_id, user_name, encrypted_password, image)
			user.save()
			return redirect(url_for('show_users'))


	@app.route('/users/<id>/edit', methods=['GET', 'POST'])
	def edit_user(id):
		user = User.get_by_id(id)
		if request.method == 'GET': 
			return render_template('user/edit.html', logged_user=User.get_logged_user())
		else:
			user_name = request.form['user_name']
			password = request.form['password']
			role_id = request.form['role_id']
			image = user.image
			new_user = User(role_id, user_name, password, image)
			new_user.save()
			return redirect(url_for('show_users'))


	@app.route('/users/search', methods=['GET', 'POST'])
	def search_user():
		if request.method == 'GET': 
			return render_template('user/search.html', logged_user=User.get_logged_user())
		else:
			value = request.form['search_value']
			res = User.search(value)
			num_rows =  len(res)
			return render_template('user/search-results.html', value=value, results=res, 
					num_rows=num_rows, logged_user=User.get_logged_user())
		

	@app.route(f'/upload-image', methods=['GET', 'POST'])
	def upload_image():
		logged_user = User.get_logged_user_basic()
		if request.method == 'GET': 
			return render_template('user/upload-image.html', logged_user=logged_user)
		else:
			user_id = logged_user.user_id
			uploader = FileUploader()
			image = uploader.upload_image('image', UPLOAD_DIR_IMGS)
			logged_user.update_image(image, user_id)
			return redirect(url_for('get_user_data'))`
}


func (mvc * MvcController) RoleController()  string {
return ``+mvcImport.ImportForRoleController("mvc")+`

class RoleController:

	@app.route('/roles', methods=['GET'])
	def show_roles():
		roles = Role.get_all()
		num_rows = len(roles)
		return render_template('role/show.html', roles=roles, num_rows=num_rows, logged_user=User.get_logged_user())


	@app.route('/roles/<id>/details', methods=['GET'])
	def role_details(id):
		role = Role.get_data_by_id(id)
		if role:
			return render_template('role/details.html', role=role, logged_user=User.get_logged_user())
		else:
			return render_template('error/404.html')


	@app.route('/roles/add', methods=['GET', 'POST'])
	def add_role():
		if request.method == 'GET': 
			return render_template('role/add.html', logged_user=User.get_logged_user())
		else:
			role_name = request.form['role_name']
			role = Role(role_name)
			role.save()
			return redirect(url_for('show_roles'))`
}


func (mvc *MvcController) AuthController()  string {
return ``+mvcImport.ImportForAuthController("mvc")+`

class AuthController:

	@app.route('/login', methods=['GET', 'POST'])
	def login():
		if request.method == 'GET':
			return render_template('auth/login.html')
		else:
			user_name = request.form['user_name']
			password = request.form['password']
			user = User.get_by_username(user_name)

			encrypted_password = user.password
			if PasswordHandler.check(encrypted_password, password):
				session['user_name'] = user_name
				session['password'] = encrypted_password
				return redirect(url_for('home'))
			else:
				return redirect(url_for('login'))


	@app.route('/logout', methods=['GET'])
	def logout():
		session.pop('user_name')
		return redirect(url_for('login'))


	@app.route('/home', methods=['GET'])
	def home():
		return render_template('auth/home.html', logged_user=User.get_logged_user())`
}


func (mvc *MvcController) FrontController() string {
return `from config import *
from flask import render_template

class FrontController:

	@app.route('/', methods=['GET'])
	def index():
		return render_template('front/index.html')

	@app.route('/register', methods=['GET'])
	def register():
		return render_template('front/index.html')`
}