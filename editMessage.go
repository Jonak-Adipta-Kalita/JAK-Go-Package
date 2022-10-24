package jak_go_package

import (
	"strings"
)

type EditMessage struct {
	Message string
}

func (x EditMessage) RemoveSpace() string {
	return strings.Replace(x.Message, " ", "", -1)
}

func (x EditMessage) ToLowerCase() string {
	return strings.ToLower(x.Message)
}

func (x EditMessage) ToUpperCase() string {
	return strings.ToUpper(x.Message)
}

func (x EditMessage) ToTitleCase() string {
	return strings.ToTitle(x.Message)
}
