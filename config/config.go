package config

type Config struct {
	DefaultLanguage string        `mapstructure:"DefaultLanguage"`
	Logging         LoggingConfig `mapstructure:"Logging"`
}

type LoggingConfig struct {
	Level      string `mapstructure:"Level"`
	Filename   string `mapstructure:"Filename"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxBackups int    `mapstructure:"MaxBackups"`
	MaxAge     int    `mapstructure:"MaxAge"`
	Compress   bool   `mapstructure:"Compress"`
}
