package config

// Config - structure describing the configuration.
type Config struct {
	NetServer
}

// NetServer - with http server options.
type NetServer struct {
	Host     string
	Port     int
	ShareDir string
	Type     string
}
