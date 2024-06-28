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
	ProxyModel,
	ProxyWeight,
}

var NodeFlag = []*Flag{
	ProxyPort,
}

var (
	// Port gateway server port
	Port = NewFlag("gateway-port", ":8000", "http server port")

	// ProxyPort Proxy Server Port
	ProxyPort = NewFlag("proxy-port", "http://127.0.0.1:8081", "http server port")

	// ProxyWeight proxy server weight
	ProxyWeight = NewFlag("proxy-weight", "1", "proxy weight")

	// ProxyModel proxy model
	ProxyModel = NewFlag("proxy-model", "rand", "proxy model")

	ReplicateCount = NewFlag("replicate-count", "3", "replicate count")
)
