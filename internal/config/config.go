package config

// Config - structure describing the configuration.
type Config struct {
	WebServe
}

// WebServe - with http server options.
type WebServe struct {
	Host string
	Port int
}
