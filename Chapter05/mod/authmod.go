package mod

type Auth interface {
	CheckAuth()
}

func CommandAuth(auth Auth) {
	// user.CheckAuth()
	auth.CheckAuth()
}
