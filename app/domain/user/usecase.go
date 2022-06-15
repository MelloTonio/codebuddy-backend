package user

type Usecase interface {
	GetAll() ([]User, error)
	Get(nickName string) (User, error)
}
