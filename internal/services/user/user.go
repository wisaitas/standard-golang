package user

type UserService interface {
	Read
	Create
	Update
	Delete
	Transaction
}

type userService struct {
	Read
	Create
	Update
	Delete
	Transaction
}

func NewUserService(
	read Read,
	create Create,
	update Update,
	delete Delete,
	transaction Transaction,
) UserService {
	return &userService{
		Read:        read,
		Create:      create,
		Update:      update,
		Delete:      delete,
		Transaction: transaction,
	}
}
