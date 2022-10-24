package jak_go_package

import (
	"strings"
)

type EditMessage struct {
	message string
}

func (x EditMessage) removeSpace() string {
	return strings.Replace(x.message, " ", "", -1)
}

func (x EditMessage) toLowerCase() string {
	return strings.ToLower(x.message)
}

func (x EditMessage) toUpperCase() string {
	return strings.ToUpper(x.message)
}

func (x EditMessage) toTitleCase() string {
	return strings.ToTitle(x.message)
}
