package user

type UserService interface {
	Get
	Post
	Update
	Delete
	Transaction
}

type userService struct {
	Get
	Post
	Update
	Delete
	Transaction
}

func NewUserService(
	get Get,
	post Post,
	update Update,
	delete Delete,
	transaction Transaction,
) UserService {
	return &userService{
		Get:         get,
		Post:        post,
		Update:      update,
		Delete:      delete,
		Transaction: transaction,
	}
}
