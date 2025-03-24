package db

func (u UserModel) Name() string {
	return u.Firstname + " " + u.Lastname
}
