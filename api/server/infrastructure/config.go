package infrastructure

type ServerConfig struct {
	AppName    string
	JwtAuthKey string
	DB         struct {
		Host     string
		UserName string
		Password string
		DBName   string
	}
	Port string
}
