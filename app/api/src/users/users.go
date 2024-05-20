package users;

type User struct{
	UserId int64
	Email string
	Password string
	City string
	Username string
	Admin bool
}

type UserLoginInfo struct{
	Email string `json:email`
	Password string `json:password`
}