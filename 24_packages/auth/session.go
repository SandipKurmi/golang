package auth

type UserSession struct {
	Username string
	Password string
	UserId int
	Email string
	Role string
}

func GetSession(username, password string) *UserSession {
	return &UserSession{
		Username: username,
		Password: password,
		UserId: 1,
		Email: "u@u.com",
		Role: "admin",
	}
}