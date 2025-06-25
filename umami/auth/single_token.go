package auth

type SingleTokenAuth struct {
	token string
}

func NewSingleTokenAuth(token string) Auth {
	return &SingleTokenAuth{token: token}
}

func (a *SingleTokenAuth) Get() (string, error) {
	return "Bearer " + a.token, nil
}

func (a *SingleTokenAuth) Header() string {
	return "Authorization"
}
