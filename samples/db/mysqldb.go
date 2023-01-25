package dbsamples

type MySqlDB struct {}

func (my *MySqlDB) CreateDatabase() string {
	return `
		CREATE DATABASE db_task;
		USE db_task;
	`
}

func (my *MySqlDB) CreateTables() string {
	return `
		DROP TABLE IF EXISTS roles;
		CREATE TABLE IF NOT EXISTS roles (
			role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			role_name VARCHAR(50) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		DROP TABLE IF EXISTS users;
		CREATE TABLE IF NOT EXISTS users (
			user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			user_name VARCHAR(50) NOT NULL UNIQUE,
			password VARCHAR(100) NOT NULL,
			role_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_role_user FOREIGN KEY(role_id) REFERENCES roles(role_id)
		);
		DROP TABLE IF EXISTS tasks;
		CREATE TABLE IF NOT EXISTS tasks (
			task_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			task_name VARCHAR(50) NOT NULL,
			description VARCHAR(200) NOT NULL,
			status ENUM('Completed', 'In Progress', 'Pending') DEFAULT 'Pending',
			user_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_task_user FOREIGN KEY(user_id) REFERENCES users(user_id)
		);
		`
}

func (my *MySqlDB) CreateSQLViews() string {
	return `
		DROP VIEW IF EXISTS view_user_tasks;
		CREATE VIEW view_user_tasks AS 
		SELECT t.task_id, t.task_name,
			t.description, t.status,
			t.created_at, t.updated_at,
			u.user_id AS user_id, u.user_name,
			r.role_name
		FROM tasks t
		JOIN users u ON(u.user_id = t.user_id)
		JOIN roles r ON(r.role_id = u.role_id)
		ORDER BY t.created_at DESC;
	`
}


func (my *MySqlDB) InsertRoles() string {
	return `
		INSERT INTO roles (role_name) VALUES ('admin');
		INSERT INTO roles (role_name) VALUES ('normal');
	`
}

func (my *MySqlDB) InsertUsers() string {
	return `
		INSERT INTO users (user_name, password, role_id) VALUES ('admin01', '12345678', 1);
		INSERT INTO users (user_name, password, role_id) VALUES ('user01', '12345678', 2);
		INSERT INTO users (user_name, password, role_id) VALUES ('user02', '12345678', 2);
		INSERT INTO users (user_name, password, role_id) VALUES ('user04', '12345678', 2);
		INSERT INTO users (user_name, password, role_id) VALUES ('user05', '12345678', 2);
	`
}

func (my *MySqlDB) InsertTasks() string {
	return `
		INSERT INTO tasks (task_name, description, user_id) VALUES ('Read a Book', 'Reading a programming book', 2);
		INSERT INTO tasks (task_name, description, user_id) VALUES ('Make a Lunch', 'Reading a programming book', 3);
		INSERT INTO tasks (task_name, description, user_id) VALUES ('Walk in the moon', 'Walking with my dog every day, at 6 AM', 4);
		INSERT INTO tasks (task_name, description, user_id) VALUES ('Send an email', 'Send an email to recruites', 5);
	`
}


