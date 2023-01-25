package projectgenerators

type ProjectGenerator interface {
	GenerateConfig()
	GenerateModels()
	GenerateViews()
	GenerateControllers()
	GenerateMySqlDB()
	GeneratePostgresDB()
	InstallDeps(platform string)
}