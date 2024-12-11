package interfaces

import (
	"strings"
)

type Formatter interface {
	Format(input_string string) (formatted_string string)
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (p PlainText) Format(input_string string) (formatted_string string) {
	return input_string
}

func (b Bold) Format(input_string string) (formatted_string string) {
	return strings.ToUpper(input_string)
}

func (c Code) Format(input_string string) (formatted_string string) {
	return strings.ToTitle(input_string)
}
