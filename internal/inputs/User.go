package inputs

type User struct {
	UserSignIn
	Name            string `json:"name,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty" binding:"required"`
	AccessToken     string `json:"access_token,omitempty"`
}

type UserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func (u *User) PrivateUser() User {
	return User{
		UserSignIn:  UserSignIn{Email: u.Email},
		Name:        u.Name,
		AccessToken: u.AccessToken,
	}
}
