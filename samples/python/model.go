package pythonsamples

type Model struct {}

var modelImport *AppImport


func (model *Model) UserModel()  string {
return ``+modelImport.ImportForUserModel()+`

class User(db.Model):

	__tablename__ = 'users'

	user_id = db.Column(db.Integer, primary_key=True)
	role_id = db.Column(db.Integer)
	user_name = db.Column(db.String(100))
	password = db.Column(db.String(150))
	image = db.Column(db.String(150))

	def __init__(self, role_id, user_name, password, image):
		self.user_name = user_name
		self.password = password
		self.role_id = role_id
		self.image = image
	
	def save(self):
		db.session.add(self)
		db.session.commit()

	def delete(self):
		db.session.delete(self)
		db.session.commit()

	@classmethod
	def exists(cls, user_name):
		return bool(cls.query.filter_by(user_name=user_name).first())

	@classmethod
	def get_by_id(cls, id):
		return cls.query.filter_by(user_id=id).first()

	@classmethod
	def get_by_username(cls, user_name):
		return cls.query.filter_by(user_name=user_name).first()

	@classmethod
	def get_all(cls):
		return cls.query.all()

	@classmethod
	def get_logged_user(cls):
		user_name = session['user_name']
		password = session['password']
		return cls.get_user_data(user_name, password)

	@classmethod
	def get_logged_user_basic(cls):
		user_name = session['user_name']
		password = session['password']
		return cls.query.filter_by(user_name=user_name, password=password).first()

	@classmethod
	def get_user_data(cls, user_name, password):
		with engine.connect() as conn:
			return conn.execute(text(f"SELECT * FROM view_user_data WHERE user_name = :user_name AND password = :password;"), {'user_name': user_name, 'password': password}).first()

	@classmethod
	def search(cls, value):
		with engine.connect() as conn:
			return conn.execute(text(f"SELECT * FROM view_user_data WHERE user_id = :value"+	
								f" OR user_name = :value "+
								f" OR role_name = :value"), {'value': value}).fetchall()

	@classmethod
	def get_data_by_id(cls, id):
		with engine.connect() as conn:
			return conn.execute(text("SELECT * FROM view_user_data WHERE user_id = :id;"), {'id': id}).first()

	@classmethod
	def update_image(cls, image, id):
		with engine.begin() as conn:
			conn.execute(text(f"UPDATE users SET image = :image WHERE user_id = :id;"), {'image': image,  'id': id})

	@classmethod
	def get_all_data(cls):
		with engine.connect() as conn:
			return conn.execute(text("SELECT * FROM view_user_data;")).fetchall()

	def to_json(self):
		user = self.get_data_by_id(self.user_id)
		return {
			"user_id": user.user_id,
			"user_name": user.user_name,
			"password": user.password,
			"image": user.image,
			"created_at": user.created_at,
			"updated_at": user.updated_at,
			"role_id": user.role_id,
			"role_name": user.role_name
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
	
	def save(self):
		db.session.add(self)
		db.session.commit()

	def delete(self):
		db.session.delete(self)
		db.session.commit()

	@classmethod
	def exists(cls, user_id, task_name, start_date, end_date):
		return bool(cls.query.filter_by(user_id=user_id, task_name=task_name, start_date=start_date, end_date=end_date).first())
		
	@classmethod
	def get_by_id(cls, id):
		return cls.query.filter_by(task_id=id).first()
	
	@classmethod
	def get_by_user(cls, user_id):
		return cls.query.filter_by(user_id=user_id).all()

	@classmethod
	def get_by_date(cls, start_date, end_date):
		return cls.query.filter(cls.start_date.between(start_date, end_date)).all()
	
	@classmethod
	def get_all(cls):
		return cls.query.all()

	@classmethod
	def get_by_status(cls, status):
		with engine.connect() as conn:
			return conn.execute(text(f"SELECT * FROM view_user_tasks WHERE status = :status;"), {'status': status}).fetchall()

	@classmethod
	def search(cls, value):
		with engine.connect() as conn:
			return conn.execute(text(f"SELECT * FROM view_user_tasks WHERE user_id = :value"+	
								f" OR user_name = :value"+
								f" OR role_name = :value"), {'value': value}).fetchall()

	@classmethod
	def get_data_by_id(cls, id):
		with engine.connect() as conn:
			return conn.execute(text(f"SELECT * FROM view_user_tasks WHERE task_id = :id;"), {'id': id}).first()

	@classmethod
	def get_all_data(cls):
		with engine.connect() as conn:
			return conn.execute(text("SELECT * FROM view_user_tasks;")).fetchall()

	def to_json(self):
		task = Task.get_data_by_id(self.task_id)
		return {
			"task_id": task.task_id,
			"task_name": task.task_name,
			"description": task.description,
			"start_date": task.start_date,
			"end_date": task.end_date,
			"user_id": task.user_id,
			"status": task.status,
			"created_at": task.created_at,
			"created_at": task.created_at,
			"user_name": task.user_name,
			"role_name": task.role_name
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
	def get_all(cls):
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