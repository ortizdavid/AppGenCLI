package pythonsamples

type Model struct {}

var appImport *AppImport

func (model *Model) UserModel()  string {
return `
`+appImport.ImportForUserModel()+`

class User(db.Model):

	__tablename__ = 'users'

	user_id = db.Column(db.Integer, primary_key=True)
    role_id = db.Column(db.Integer)
    email = db.Column(db.String(100))
    password = db.Column(db.String(100))

    def __init__(self, role_id, email, password):
        self.email = email
        self.password = password
        self.role_id = role_id

	`+model.SaveAndDelete()+`

	@classmethod
	def get_logged_user(cls):
		email = session['email']
		password = session['password']
		return cls.get_user_data(email, password)

	@classmethod
	def get_user_data(cls, email, password):
		return engine.execute(f"SELECT * FROM view_user_data WHERE email = '{email}' AND password='{password}';").first()

	@classmethod
	def exists(cls, email, senha) -> bool:
		return bool(cls.query.filter_by(email=email, senha=senha).first())

	@classmethod
	def get_by_id(cls, id):
		return engine.execute("SELECT * FROM view_user_data WHERE user_id = {id};").first()
	
	@classmethod
	def get_all(cls):
		return engine.execute("SELECT * FROM view_user_data;").fetchall()

	def to_json(self):
		return {
			"name":
		}
`
}



func (model *Model) TaskModel() string {
return `
class Task(db.Model):

	__tablename__ = 'tasks'

	def to_json(self):
		return {
			
		}
`
}

func (model *Model) SaveAndDelete() string {
return `
def save(self):
	db.session.add(self)
	db.session.commit()

def delete(self):
	db.session.delete(self)
	db.session.commit()
`
}