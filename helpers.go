package main

/////////////////////
// helper wrappers
/////////////////////
func info(msg string, params ...interface{}) {
	app.Logger.Infow(msg, params...)
}

func warn(msg string, params ...interface{}) {
	app.Logger.Warnw(msg, params...)
}

func alert(msg string, params ...interface{}) {
	app.Logger.Errorw(msg, params...)
}
