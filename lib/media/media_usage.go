package media

type Usage int

const (
	Undefined Usage = iota
	UserAvatar
	GroupAvatar
)

func Parse(str string) Usage {
	switch str {
	case "userAvatar":
		return UserAvatar
	case "groupAvatar":
		return GroupAvatar
	default:
		return Undefined
	}
}
