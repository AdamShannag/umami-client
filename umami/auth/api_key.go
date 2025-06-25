package auth

type ApiKeyAuth struct {
	key string
}

func NewApiKeyAuth(key string) Auth {
	return &ApiKeyAuth{key: key}
}

func (a *ApiKeyAuth) Get() (string, error) {
	return a.key, nil
}

func (a *ApiKeyAuth) Header() string {
	return "x-umami-api-key"
}
