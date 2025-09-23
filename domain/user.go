package domain

// model or entity -> existence
type User struct {
	ID           int    `db:"id" json:"id"`
	FullName     string `db:"full_name" json:"fullName"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"password" json:"password"`
	IsShopeOwner bool   `db:"is_shope_owner" json:"isShopeOwner"`
}
