package config

type ServerConfiguration struct {
	Port int
}

type Configuration struct {
	Server ServerConfiguration
	Upload string
}
