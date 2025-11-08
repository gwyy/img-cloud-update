package config

type Config struct {
	Mode    string  `yaml:"mode" mapstructure:"mode"`
	Version string  `yaml:"version" mapstructure:"version"`
	Gin     Gin     `yaml:"gin" mapstructure:"gin"`
	BboltDB BboltDB `yaml:"bbolt_db" mapstructure:"bbolt_db"`
	Log     Log     `yaml:"log" mapstructure:"log"`
}

type BboltDB struct {
	Path      string `yaml:"path" mapstructure:"path"`
	TableName string `yaml:"table_name" mapstructure:"table_name"`
}

type Gin struct {
	Name         string `yaml:"name" mapstructure:"name"`
	Host         string `yaml:"host" mapstructure:"host"`
	Prefix       string `yaml:"prefix" mapstructure:"prefix"`
	IP           string `yaml:"ip" mapstructure:"ip"`
	Port         int    `yaml:"port" mapstructure:"port"`
	Timeout      int    `yaml:"timeout" mapstructure:"timeout"`
	Pprof        bool   `yaml:"pprof" mapstructure:"pprof"`
	ReadTimeout  string `yaml:"readtimeout" mapstructure:"readtimeout"`
	WriteTimeout string `yaml:"writetimeout" mapstructure:"writetimeout"`
}

type Log struct {
	Level      int    `yaml:"level" mapstructure:"level"`
	Path       string `yaml:"path" mapstructure:"path"`
	MaxSize    int    `yaml:"max-size" mapstructure:"max-size"`
	MaxAge     int    `yaml:"max-age" mapstructure:"max-age"`
	MaxBackups int    `yaml:"max-backups" mapstructure:"max-backups"`
	Compress   bool   `yaml:"compress" mapstructure:"compress"`
}
