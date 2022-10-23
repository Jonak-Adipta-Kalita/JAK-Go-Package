package editMessage

import (
	"strings"
)

type Message struct {
	message string
}

func (x Message) removeSpace() string {
	return strings.Replace(x.message, " ", "", -1)
}

func (x Message) toLowerCase() string {
	return strings.ToLower(x.message)
}

func (x Message) toUpperCase() string {
	return strings.ToUpper(x.message)
}

func (x Message) toTitleCase() string {
	return strings.ToTitle(x.message)
}
