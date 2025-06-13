package user

type UserService interface {
	Get
	Post
	Update
	Delete
}

type userService struct {
	Get
	Post
	Update
	Delete
}

func NewUserService(
	get Get,
	post Post,
	update Update,
	delete Delete,
) UserService {
	return &userService{
		Get:    get,
		Post:   post,
		Update: update,
		Delete: delete,
	}
}
