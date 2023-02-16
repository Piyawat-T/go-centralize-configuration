package bootstrap

type Application struct {
	Env      *Env
	Database Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Database = NewMySQLDatabase(app.Env)
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	app.Database.CloseMySQLDatabase()
}
