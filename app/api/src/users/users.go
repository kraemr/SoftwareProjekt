package users;

type User struct{
	Email string
	Password string
	City string
	UserId int64
	Username string
	Admin bool
}

type UserLoginInfo struct{
	Email string `json:email`
	Password string `json:password`
}