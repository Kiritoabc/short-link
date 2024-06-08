package config

type Flag struct {
	Name        string // flag name
	Value       string // flag value
	Description string // flag Description
}

func NewFlag(name, value, description string) *Flag {
	return &Flag{
		Name:        name,
		Value:       value,
		Description: description,
	}
}

var Flags = []*Flag{}
