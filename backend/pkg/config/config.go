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

var (
	// Port server port
	Port = NewFlag("server-port", ":8080", "http server port")
	// DatabaseDriver  database driver
	DatabaseDriver = NewFlag("use-database-driver", "mysql", "what database driver your want to use,now support [mysql]")
	// EnableSwagger swagger enabled
	EnableSwagger = NewFlag("enable-swagger", "true", "wether swagger eanbled")
	// SnowFlakeNode snowflake
	SnowFlakeNode = NewFlag("snowfalke node", "1", "snowflake node num")

	// SLDatabaseHost
	// SLDatabasePort
	// SLDatabaseUsername
	// SLDatabasePassword
	// SLDatabaseScheme
	// SLDatabaseArgs
	SLDatabaseHost     = NewFlag("short-link-database-host", "127.0.0.1", "short link database connect host")
	SLDatabasePort     = NewFlag("short-link-database-port", "3306", "short link database connect port")
	SLDatabaseUsername = NewFlag("short-link-database-username", "root", "short link database connect username")
	SLDatabasePassword = NewFlag("short-link-database-password", "123456", "short link database connect password")
	SLDatabaseScheme   = NewFlag("short-link-database-scheme", "shortlink", "short link which database you select")
	SLDatabaseArgs     = NewFlag("short-link-database-args", "charset=utf8mb4&loc=Local&parseTime=true", "short link database connect args")
)
