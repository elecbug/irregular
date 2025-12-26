package date

// Config holds configuration options for date parsing.
type Config struct {
	BaseYear int
}

// DefaultConfig returns the default configuration for date parsing.
func DefaultConfig() *Config {
	return &Config{
		BaseYear: 1970,
	}
}
