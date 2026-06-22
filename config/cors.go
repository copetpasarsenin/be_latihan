package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://127.0.0.1:5173",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
