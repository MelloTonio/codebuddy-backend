package user

type Repository interface {
	GetAll() ([]User, error)
	Get(nickname string) (User, error)
}
