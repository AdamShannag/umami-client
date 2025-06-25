package auth

type Auth interface {
	Header() string
	Get() (string, error)
}
