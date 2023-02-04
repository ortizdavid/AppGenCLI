package dbsamples

type SampleDB interface {
	CreateDatabase() string
	CreateTables() string
	CreateSQLViews() string
	InsertRoles() string
	InsertUsers() string
	InsertTasks() string
	GetDatabaseScript() string
}