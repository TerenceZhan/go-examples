package Users

type Users struct {
	id       int
	email    string
	password string
}

var u1 = Users{id: 1, email: "test@terence.com", password: "xxx"}

func existsUser(user Users) bool {
	return user.id == u1.id
}
