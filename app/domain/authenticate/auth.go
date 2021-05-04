package access

type Service interface {
	Authenticate(Credential) (Description, error)
}
