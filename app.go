package antgolang

import "strings"

type AntApp struct {
	Authorization string
	Exchange      string
}

func New(exchange string) AntApp {
	app := AntApp{
		Authorization: "",
		Exchange:      strings.ToUpper(exchange),
	}
	return app
}

func (app *AntApp) SetToken(userId, token string) {
	app.Authorization = "Bearer " + userId + " " + token
}
