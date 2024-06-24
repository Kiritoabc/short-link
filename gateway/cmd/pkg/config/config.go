package config

type Flag struct {
	Name        string // flag name
	Value       string // flag value
	Description string // flag description
}

func NewFlag(name, value, description string) *Flag {
	return &Flag{
		Name:        name,
		Value:       value,
		Description: description,
	}
}

var Flags = []*Flag{
	Port,
	ProxyPort,
}

var (
	// Port gateway server port
	Port = NewFlag("gateway-port", ":8000", "http server port")

	// ProxyPort Proxy Server Port
	ProxyPort = NewFlag("proxy-port", "http://127.0.0.1:8081", "http server port")

	// ProxyModel proxy model
	ProxyModel = NewFlag("proxy-model", "rand", "proxy model")
)
