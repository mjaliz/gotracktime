package inputs

type User struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password,omitempty" binding:"required"`
	PasswordConfirm string `json:"password_confirm,omitempty" binding:"required"`
	AccessToken     string `json:"access_token,omitempty"`
}

func (u *User) PrivateUser() User {
	return User{
		Name:        u.Name,
		Email:       u.Email,
		AccessToken: u.AccessToken,
	}
}
