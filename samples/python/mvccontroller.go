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

	@app.route('/users', methods=['GET'])
	def show_users():
		users = User.get_all()
		num_rows = len(users)
		return render_template('user/show.html', users=users, num_rows=num_rows, logged_user=User.get_logged_user())

	@app.route('/users/<id>/details', methods=['GET'])
	def user_details(id):
		user = User.get_data_by_id(id)
		if user:
			return render_template('user/details.html', user=user, logged_user=User.get_logged_user())
		else:
			return render_template('errorr/404.html')

	@app.route(f'/dados-pessoais', methods=['GET'])
	def get_user_data():
		logged_user = User.obter_logged_user()
		data = User.get_by_id(logged_user.user_id)
		return render_template('user/user-data.html', data=data, logged_user=logged_user)

	@app.route('/users/add', methods=['GET', 'POST'])
	def add_user():
		if request.method == 'GET': 
			return render_template('user/add.html', logged_user=User.get_logged_user())
		else:
			user_name = request.form['user_name']
			password = request.form['password']
			role_id = request.form['role_id']
			file = request.form['image']
			image = secure_filename(file.filename)
			file.save(os.path.join(UPLOAD_DIR_IMGS, image))
			user = User(role_id, user_name, password, image)
			user.save()
			return redirect(url_for('show_users'))

	@app.user('/users/<id>/edit', methods=['GET', 'POST'])
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
			return render_template('user/search.html',logged_user=User.get_logged_user())
		else:
			value = request.form['search_value']
			res = User.search(value)
			num_rows =  len(res)
			return render_template('user/search-results.html', value=value, results=res, 
					num_rows=num_rows, logged_user=User.get_logged_user())`
}


func (mvc * MvcController) RoleController()  string {
return ``+mvcImport.ImportForRoleController()+`

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


func (mvc * MvcController) AuthController()  string {
return ``+mvcImport.ImportForAuthController()+`

class AuthController:

	@app.route('/', methods=['GET'])
	@app.route('/login', methods=['GET', 'POST'])
	def login():
		if request.method == 'GET':
			return render_template('auth/login.html')
		else:
			user_name = request.form['user_name']
			password = request.form['password']
			if(User.exists(user_name, password)):
				session['user_name'] = user_name
				session['password'] = password
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