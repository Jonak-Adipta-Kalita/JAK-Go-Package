package jak_go_package

import (
	"strings"
)

type editMessage struct {
	message string
}

func (x editMessage) removeSpace() string {
	return strings.Replace(x.message, " ", "", -1)
}

func (x editMessage) toLowerCase() string {
	return strings.ToLower(x.message)
}

func (x editMessage) toUpperCase() string {
	return strings.ToUpper(x.message)
}

func (x editMessage) toTitleCase() string {
	return strings.ToTitle(x.message)
}
