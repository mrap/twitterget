package user

type User struct {
	ScreenName string
}

func NewUser(screenName string) *User {
	return &User{
		ScreenName: PlainScreenName(screenName),
	}
}

func PlainScreenName(name string) string {
	for i, c := range name {
		switch string(c) {
		case "@", " ":
			continue
		default:
			return name[i:len(name)]
		}
	}
	return ""
}
