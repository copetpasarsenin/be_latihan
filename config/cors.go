package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://127.0.0.1:5173",
	"belatihan-production.up.railway.app",
	"https://my-fe-gamma.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
