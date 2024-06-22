package media

import monify "monify/protobuf/gen/go"

func Parse(str string) monify.Usage {
	println(str)
	switch str {
	case "userAvatar":
		return monify.Usage_UserAvatar
	case "groupAvatar":
		return monify.Usage_GroupAvatar
	default:
		return monify.Usage_Undefined
	}
}
