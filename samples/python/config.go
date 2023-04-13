package pythonsamples

type ConfigPy struct {}


func (conf *ConfigPy) CreateConfig(db string) string   {

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

APP_PORT = "5000"
APP_ROOT = f"localhost:{APP_PORT}"
API_ROOT = "/api"

UPLOAD_DIR_IMGS = "static/uploads/imgs"
UPLOAD_DIR_DOCS = "static/uploads/docs"

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = DB_URI
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.secret_key = 'my-app'

db = SQLAlchemy()
db.init_app(app)

engine = create_engine(DB_URI)`
}