package repo

type User struct {
	ID           int    `json:"id"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopeOwner bool   `json:"isShopeOwner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	UserList []User
}

// Constructor
func NewUserRepo() UserRepo {
	return &userRepo{}
}

// Create user
func (r *userRepo) Create(user User) (*User, error) {
	// auto ID generate
	user.ID = len(r.UserList) + 1

	r.UserList = append(r.UserList, user)
	return &user, nil
}

// Find user by email & password
func (r *userRepo) Find(email, pass string) (*User, error) {
	for i := range r.UserList {
		if r.UserList[i].Email == email && r.UserList[i].Password == pass {
			return &r.UserList[i], nil
		}
	}
	return nil, nil
}
