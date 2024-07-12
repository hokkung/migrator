package config

const AppPrefix = "APP"

type Configuration struct {
	AppProperties   AppProperties
	MysqlProperties MysqlProperties
}

// NewConfiguration returns new config instance
func NewConfiguration(
	app AppProperties,
	mysql MysqlProperties,
) *Configuration {
	return &Configuration{
		AppProperties:   app,
		MysqlProperties: mysql,
	}
}
