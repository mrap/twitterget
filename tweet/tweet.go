package tweet

import (
	"github.com/mrap/twitterget/user"
)

type Tweet struct {
	ID     string
	Text   string
	Author user.User
}
