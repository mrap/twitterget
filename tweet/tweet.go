package tweet

import (
	"github.com/mrap/twitterget/user"
)

type Tweet struct {
	Text   string
	Author user.User
}
