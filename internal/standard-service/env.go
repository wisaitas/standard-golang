package standardservice

var ENV struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Redis    `mapstructure:"redis"`
}

type Server struct {
	Port        int    `mapstructure:"port" defaultValue:"8000"`
	MaxFileSize int64  `mapstructure:"max_file_size" defaultValue:"5"`
	JwtSecret   string `mapstructure:"jwt_secret" defaultValue:"secret"`
}

type Database struct {
	Host     string `mapstructure:"host" defaultValue:"localhost"`
	Port     int    `mapstructure:"port" defaultValue:"9000"`
	User     string `mapstructure:"user" defaultValue:"postgres"`
	Password string `mapstructure:"password" defaultValue:"root"`
	Name     string `mapstructure:"name" defaultValue:"postgres"`
	Driver   string `mapstructure:"driver" defaultValue:"postgres"`
}

type Redis struct {
	Host     string `mapstructure:"host" defaultValue:"localhost"`
	Port     int    `mapstructure:"port" defaultValue:"9001"`
	Password string `mapstructure:"password" defaultValue:""`
}
