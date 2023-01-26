package pythonsamples


type ConfigPy struct {}


func (conf *ConfigPy) GetConfig(db string) string   {

	var rdms, dbUser, dbPort string
	var dbPassword = ""
	var dbHost = "localhost"
	var dbName = "db_task"

	switch db {
	case "mysql":
		rdms = "mysql"
		dbUser = "root"
		dbPort = "3306"
	case "postgres":
		rdms = "postgresql"
		dbUser = "postgres"
		dbPort = "5432"
	default:
		rdms = ""	
	}

	return `

		RDMS = "`+rdms+`"
		DB_USER = "`+dbUser+`"
		DB_PASSWORD = "`+dbPassword+`"
		DB_HOST = "`+dbHost+`"
		DB_PORT = "`+dbPort+`"
		DB_NAME = "`+dbName+`"
		DB_URI = f"{RDMS}://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}"

		app = Flask(__name__)
		app.config['SQLALCHEMY_DATABASE_URI'] = DB_URI
		app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
		app.secret_key = 'my-app'

		db = SQLAlchemy()
		db.init_app(app)

		engine = create_engine(DB_URI)
	`
}