package repositories

type UserRepository interface {
	GetUserNames() []string
}

type userRepository struct {
	// user_name string
}

func NewUserRepository() UserRepository {
	// return &userRepository{user_name: "デバッグ用"}
	return &userRepository{}
}

func (r *userRepository) GetUserNames() []string {
	return []string{"Alice", "Bob", "Carlie"}
	// return []string{r.user_name}
}
