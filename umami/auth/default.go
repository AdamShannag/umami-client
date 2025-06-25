package auth

type DefaultAuth struct {
}

func NewDefaultAuth() Auth {
	return &DefaultAuth{}
}

func (a *DefaultAuth) Get() (string, error) {
	return "", nil
}

func (a *DefaultAuth) Header() string {
	return ""
}
