package access

type Service interface {
	Authenticate(Credential) (string, error)
}
