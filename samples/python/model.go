package pythonsamples

type Model struct {}

var modelImport *AppImport

func (model *Model) UserModel()  string {
return ``+modelImport.ImportForUserModel()+`

class User(db.Model):

	__tablename__ = 'users'

	user_id = db.Column(db.Integer, primary_key=True)
	role_id = db.Column(db.Integer)
	user_name = db.Column(db.String(150))
	password = db.Column(db.String(100))
	image = db.Column(db.String(100))

	def __init__(self, role_id, user_name, password, image):
		self.user_name = user_name
		self.password = password
		self.role_id = role_id
		self.image = image
	`+model.SaveAndDelete()+`
	@classmethod
	def exists(cls, user_name, password):
		return bool(cls.query.filter_by(user_name=user_name, password=password).first())

	@classmethod
	def get_by_id(cls, id):
		return cls.query.filter_by(user_id=id).first()

	@classmethod
	def get_all(cls, id):
		return cls.query.fetchall()

	@classmethod
	def get_logged_user(cls):
		user_name = session['user_name']
		password = session['password']
		return cls.get_user_data(user_name, password)

	@classmethod
	def get_user_data(cls, user_name, password):
		return engine.execute(f"SELECT * FROM view_user_data WHERE user_name = '{user_name}' AND password='{password}';").first()

	@classmethod
	def get_data_by_id(cls, id):
		return engine.execute("SELECT * FROM view_user_data WHERE user_id = {id};").first()

	@classmethod
	def get_all_data(cls):
		return engine.execute("SELECT * FROM view_user_data;").all()

	def to_json(self):
		return {
			"user_id": self.user_id,
			"user_name": self.user_name,
			"password": self.password,
			"role_id": self.role_id,
			"image": self.image
		}`
}

func (model *Model) TaskModel()  string {
return ``+modelImport.ImportForAllModels()+`

class Task(db.Model):

	__tablename__ = 'tasks'

	task_id = db.Column(db.Integer, primary_key=True)
	user_id = db.Column(db.Integer)
	task_name = db.Column(db.String(100))
	start_date = db.Column(db.Date)
	end_date = db.Column(db.Date)
	description = db.Column(db.String(300))

	def __init__(self, user_id, task_name, start_date, end_date, description):
		self.task_name = task_name
		self.description = description
		self.user_id = user_id
		self.start_date = start_date
		self.end_date = end_date
	`+model.SaveAndDelete()+`
	@classmethod
	def exists(cls, user_id, task_name):
		return bool(cls.query.filter_by(user_id=user_id, task_name=task_name).first())

	@classmethod
	def get_by_id(cls, id):
		return cls.query.filter_by(task_id=id).first()

	@classmethod
	def get_all(cls, id):
		return cls.query.fetchall()

	@classmethod
	def get_data_by_id(cls, id):
		return engine.execute("SELECT * FROM view_user_tasks WHERE user_id = {id};").first()

	@classmethod
	def get_all_data(cls):
		return engine.execute("SELECT * FROM view_user_tasks;").all()

	def to_json(self):
		return {
			"task_id": self.task_id,
			"user_id": self.user_id,
			"task_name": self.task_name,
			"start_date": self.start_date,
			"end_date": self.end_date,
			"description": self.description
		}`
}

func (model *Model) RoleModel()  string {
return ``+modelImport.ImportForAllModels()+`

class Role(db.Model):

	__tablename__ = 'roles'

	role_id = db.Column(db.Integer, primary_key=True)
	role_name = db.Column(db.String(100))

	def __init__(self, role_name):
		self.role_name = role_name

	`+model.SaveAndDelete()+`
	
	@classmethod
	def exists(cls, role_name):
		return bool(cls.query.filter_by(role_name=role_name).first())

	@classmethod
	def get_by_id(cls, id):
		return cls.query.filter_by(role_id=id).first()

	@classmethod
	def get_all(cls, id):
		return cls.query.all()

	def to_json(self):
		return {
			"role_id": self.role_id,
			"role_name": self.role_name
		}`
}

func (model *Model) SaveAndDelete() string {
return `
	def save(self):
		db.session.add(self)
		db.session.commit()

	def delete(self):
		db.session.delete(self)
		db.session.commit()`
}