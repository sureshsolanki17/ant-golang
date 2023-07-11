package antgolang

type AntApp struct {
	Authorization string
}

func New() AntApp {
	return AntApp{}
}

func (an *AntApp) SetToken(userId, token string) {
	Authorization = "Bearer " + userId + " " + token
	an.Authorization = Authorization
}
